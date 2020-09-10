package job

import (
	"github.com/smartcontractkit/chainlink/core/services/offchainreporting"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

type (
	Type string

	Spec interface {
		JobID() *models.ID
		JobType() Type
		TaskDAG() pipeline.TaskDAG
	}

	Service interface {
		Start() error
		Stop() error
	}

	// Job must conform to the Spec interface
	Job struct {
		ID                            int32 `gorm: "primarykey"`
		OffchainreportingOracleSpecID int32
		OffchainreportingOracleSpec   offchainreporting.OracleSpec
	}
)

// NOTE: Since the only possible type of job for now is offchainreporting,
// these implementations can be simplistic
func (j *Job) JobID() *models.ID {
	return nil
}

func (j *Job) JobType() Type {
	return "offchainreporting_oracle"
}

func (j *Job) TaskDAG() pipeline.TaskDAG {
	return j.OffchainreportingOracleSpec
}
