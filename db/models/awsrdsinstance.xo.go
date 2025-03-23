package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AwsRdsInstance represents a row from 'public.aws_rds_instances'.
type AwsRdsInstance struct {
	ID                             string         `json:"id" gorm:"column:id"`                                                                 // id
	Regioncode                     AwsRegion      `json:"regioncode" gorm:"column:regioncode"`                                                 // regioncode
	Instancetype                   string         `json:"instancetype" gorm:"column:instancetype"`                                             // instancetype
	Instancefamily                 string         `json:"instancefamily" gorm:"column:instancefamily"`                                         // instancefamily
	Vcpu                           StringInt      `json:"vcpu" gorm:"column:vcpu"`                                                             // vcpu
	Physicalprocessor              string         `json:"physicalprocessor" gorm:"column:physicalprocessor"`                                   // physicalprocessor
	Clockspeed                     string         `json:"clockspeed" gorm:"column:clockspeed"`                                                 // clockspeed
	Memory                         StringFloat    `json:"memory" gorm:"column:memory"`                                                         // memory
	Storage                        string         `json:"storage" gorm:"column:storage"`                                                       // storage
	Networkperformance             string         `json:"networkperformance" gorm:"column:networkperformance"`                                 // networkperformance
	Databaseengine                 DatabaseEngine `json:"databaseengine" gorm:"column:databaseengine"`                                         // databaseengine
	Deploymentoption               string         `json:"deploymentoption" gorm:"column:deploymentoption"`                                     // deploymentoption
	Dedicatedebsthroughput         string         `json:"dedicatedebsthroughput" gorm:"column:dedicatedebsthroughput"`                         // dedicatedebsthroughput
	Ondemandprice                  StringFloat    `json:"ondemandprice" gorm:"column:ondemandprice"`                                           // ondemandprice
	OneYearReservedStandardPrice   StringFloat    `json:"one_year_reserved_standard_price" gorm:"column:one_year_reserved_standard_price"`     // one_year_reserved_standard_price
	ThreeYearReservedStandardPrice StringFloat    `json:"three_year_reserved_standard_price" gorm:"column:three_year_reserved_standard_price"` // three_year_reserved_standard_price
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [AwsRdsInstance] exists in the database.
func (ari *AwsRdsInstance) Exists() bool {
	return ari._exists
}

// Deleted returns true when the [AwsRdsInstance] has been marked for deletion
// from the database.
func (ari *AwsRdsInstance) Deleted() bool {
	return ari._deleted
}

// Insert inserts the [AwsRdsInstance] to the database.
func (ari *AwsRdsInstance) Insert(ctx context.Context, db DB) error {
	switch {
	case ari._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ari._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.aws_rds_instances (` +
		`id, regioncode, instancetype, instancefamily, vcpu, physicalprocessor, clockspeed, memory, storage, networkperformance, databaseengine, deploymentoption, dedicatedebsthroughput, ondemandprice, one_year_reserved_standard_price, three_year_reserved_standard_price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`)`
	// run
	logf(sqlstr, ari.ID, ari.Regioncode, ari.Instancetype, ari.Instancefamily, ari.Vcpu, ari.Physicalprocessor, ari.Clockspeed, ari.Memory, ari.Storage, ari.Networkperformance, ari.Databaseengine, ari.Deploymentoption, ari.Dedicatedebsthroughput, ari.Ondemandprice, ari.OneYearReservedStandardPrice, ari.ThreeYearReservedStandardPrice)
	if _, err := db.ExecContext(ctx, sqlstr, ari.ID, ari.Regioncode, ari.Instancetype, ari.Instancefamily, ari.Vcpu, ari.Physicalprocessor, ari.Clockspeed, ari.Memory, ari.Storage, ari.Networkperformance, ari.Databaseengine, ari.Deploymentoption, ari.Dedicatedebsthroughput, ari.Ondemandprice, ari.OneYearReservedStandardPrice, ari.ThreeYearReservedStandardPrice); err != nil {
		return logerror(err)
	}
	// set exists
	ari._exists = true
	return nil
}

// Update updates a [AwsRdsInstance] in the database.
func (ari *AwsRdsInstance) Update(ctx context.Context, db DB) error {
	switch {
	case !ari._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ari._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.aws_rds_instances SET ` +
		`regioncode = $1, instancetype = $2, instancefamily = $3, vcpu = $4, physicalprocessor = $5, clockspeed = $6, memory = $7, storage = $8, networkperformance = $9, databaseengine = $10, deploymentoption = $11, dedicatedebsthroughput = $12, ondemandprice = $13, one_year_reserved_standard_price = $14, three_year_reserved_standard_price = $15 ` +
		`WHERE id = $16`
	// run
	logf(sqlstr, ari.Regioncode, ari.Instancetype, ari.Instancefamily, ari.Vcpu, ari.Physicalprocessor, ari.Clockspeed, ari.Memory, ari.Storage, ari.Networkperformance, ari.Databaseengine, ari.Deploymentoption, ari.Dedicatedebsthroughput, ari.Ondemandprice, ari.OneYearReservedStandardPrice, ari.ThreeYearReservedStandardPrice, ari.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ari.Regioncode, ari.Instancetype, ari.Instancefamily, ari.Vcpu, ari.Physicalprocessor, ari.Clockspeed, ari.Memory, ari.Storage, ari.Networkperformance, ari.Databaseengine, ari.Deploymentoption, ari.Dedicatedebsthroughput, ari.Ondemandprice, ari.OneYearReservedStandardPrice, ari.ThreeYearReservedStandardPrice, ari.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [AwsRdsInstance] to the database.
func (ari *AwsRdsInstance) Save(ctx context.Context, db DB) error {
	if ari.Exists() {
		return ari.Update(ctx, db)
	}
	return ari.Insert(ctx, db)
}

// Upsert performs an upsert for [AwsRdsInstance].
func (ari *AwsRdsInstance) Upsert(ctx context.Context, db DB) error {
	switch {
	case ari._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.aws_rds_instances (` +
		`id, regioncode, instancetype, instancefamily, vcpu, physicalprocessor, clockspeed, memory, storage, networkperformance, databaseengine, deploymentoption, dedicatedebsthroughput, ondemandprice, one_year_reserved_standard_price, three_year_reserved_standard_price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`regioncode = EXCLUDED.regioncode, instancetype = EXCLUDED.instancetype, instancefamily = EXCLUDED.instancefamily, vcpu = EXCLUDED.vcpu, physicalprocessor = EXCLUDED.physicalprocessor, clockspeed = EXCLUDED.clockspeed, memory = EXCLUDED.memory, storage = EXCLUDED.storage, networkperformance = EXCLUDED.networkperformance, databaseengine = EXCLUDED.databaseengine, deploymentoption = EXCLUDED.deploymentoption, dedicatedebsthroughput = EXCLUDED.dedicatedebsthroughput, ondemandprice = EXCLUDED.ondemandprice, one_year_reserved_standard_price = EXCLUDED.one_year_reserved_standard_price, three_year_reserved_standard_price = EXCLUDED.three_year_reserved_standard_price `
	// run
	logf(sqlstr, ari.ID, ari.Regioncode, ari.Instancetype, ari.Instancefamily, ari.Vcpu, ari.Physicalprocessor, ari.Clockspeed, ari.Memory, ari.Storage, ari.Networkperformance, ari.Databaseengine, ari.Deploymentoption, ari.Dedicatedebsthroughput, ari.Ondemandprice, ari.OneYearReservedStandardPrice, ari.ThreeYearReservedStandardPrice)
	if _, err := db.ExecContext(ctx, sqlstr, ari.ID, ari.Regioncode, ari.Instancetype, ari.Instancefamily, ari.Vcpu, ari.Physicalprocessor, ari.Clockspeed, ari.Memory, ari.Storage, ari.Networkperformance, ari.Databaseengine, ari.Deploymentoption, ari.Dedicatedebsthroughput, ari.Ondemandprice, ari.OneYearReservedStandardPrice, ari.ThreeYearReservedStandardPrice); err != nil {
		return logerror(err)
	}
	// set exists
	ari._exists = true
	return nil
}

// Delete deletes the [AwsRdsInstance] from the database.
func (ari *AwsRdsInstance) Delete(ctx context.Context, db DB) error {
	switch {
	case !ari._exists: // doesn't exist
		return nil
	case ari._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.aws_rds_instances ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, ari.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ari.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ari._deleted = true
	return nil
}

// AwsRdsInstanceByID retrieves a row from 'public.aws_rds_instances' as a [AwsRdsInstance].
//
// Generated from index 'aws_rds_instances_pkey'.
func AwsRdsInstanceByID(ctx context.Context, db DB, id string) (*AwsRdsInstance, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, regioncode, instancetype, instancefamily, vcpu, physicalprocessor, clockspeed, memory, storage, networkperformance, databaseengine, deploymentoption, dedicatedebsthroughput, ondemandprice, one_year_reserved_standard_price, three_year_reserved_standard_price ` +
		`FROM public.aws_rds_instances ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	ari := AwsRdsInstance{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ari.ID, &ari.Regioncode, &ari.Instancetype, &ari.Instancefamily, &ari.Vcpu, &ari.Physicalprocessor, &ari.Clockspeed, &ari.Memory, &ari.Storage, &ari.Networkperformance, &ari.Databaseengine, &ari.Deploymentoption, &ari.Dedicatedebsthroughput, &ari.Ondemandprice, &ari.OneYearReservedStandardPrice, &ari.ThreeYearReservedStandardPrice); err != nil {
		return nil, logerror(err)
	}
	return &ari, nil
}
