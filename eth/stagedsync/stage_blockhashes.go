package stagedsync

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/common/etl"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/eth/stagedsync/stages"
	"github.com/ledgerwatch/turbo-geth/ethdb"
)

func extractHeaders(k []byte, v []byte, next etl.ExtractNextFunc) error {
	// We only want to extract entries composed by Block Number + Header Hash
	if len(k) != 40 {
		return nil
	}
	return next(k, common.CopyBytes(k[8:]), common.CopyBytes(k[:8]))
}

type BlockHashesCfg struct {
	db     ethdb.RwKV
	tmpDir string
}

func StageBlockHashesCfg(db ethdb.RwKV, tmpDir string) BlockHashesCfg {
	return BlockHashesCfg{
		db:     db,
		tmpDir: tmpDir,
	}
}

func SpawnBlockHashStage(s *StageState, tx ethdb.RwTx, cfg BlockHashesCfg, quit <-chan struct{}) error {
	useExternalTx := tx != nil
	if !useExternalTx {
		var err error
		tx, err = cfg.db.BeginRw(context.Background())
		if err != nil {
			return err
		}
		defer tx.Rollback()
	}
	headNumber, err := stages.GetStageProgress(tx, stages.Headers)
	if err != nil {
		return fmt.Errorf("getting headers progress: %w", err)
	}
	headHash := rawdb.ReadHeaderByNumber(tx, headNumber).Hash()
	if s.BlockNumber == headNumber {
		s.Done()
		return nil
	}

	startKey := make([]byte, 8)
	binary.BigEndian.PutUint64(startKey, s.BlockNumber)
	endKey := dbutils.HeaderKey(headNumber, headHash) // Make sure we stop at head

	//todo do we need non canonical headers ?
	logPrefix := s.state.LogPrefix()
	if err := etl.Transform(
		logPrefix,
		tx,
		dbutils.HeadersBucket,
		dbutils.HeaderNumberBucket,
		cfg.tmpDir,
		extractHeaders,
		etl.IdentityLoadFunc,
		etl.TransformArgs{
			ExtractStartKey: startKey,
			ExtractEndKey:   endKey,
			Quit:            quit,
		},
	); err != nil {
		return err
	}
	if err := s.DoneAndUpdate(tx, headNumber); err != nil {
		return err
	}
	if !useExternalTx {
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

func UnwindBlockHashStage(u *UnwindState, s *StageState, tx ethdb.RwTx, cfg BlockHashesCfg) error {
	useExternalTx := tx != nil
	if !useExternalTx {
		var err error
		tx, err = cfg.db.BeginRw(context.Background())
		if err != nil {
			return err
		}
		defer tx.Rollback()
	}

	err := u.Done(tx)
	logPrefix := s.state.LogPrefix()
	if err != nil {
		return fmt.Errorf("%s: reset: %v", logPrefix, err)
	}
	if !useExternalTx {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("%s: failed to write db commit: %v", logPrefix, err)
		}
	}
	return nil
}
