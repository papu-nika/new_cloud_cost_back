package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AwsEc2Inctance represents a row from 'public.aws_ec2_inctances'.
type AwsEc2Inctance struct {
	ID                                string      `json:"id" gorm:"column:id"`                                                                       // id
	Regioncode                        AwsRegion   `json:"regioncode" gorm:"column:regioncode"`                                                       // regioncode
	Instancetype                      string      `json:"instancetype" gorm:"column:instancetype"`                                                   // instancetype
	Instancefamily                    string      `json:"instancefamily" gorm:"column:instancefamily"`                                               // instancefamily
	Vcpu                              StringInt   `json:"vcpu" gorm:"column:vcpu"`                                                                   // vcpu
	Physicalprocessor                 string      `json:"physicalprocessor" gorm:"column:physicalprocessor"`                                         // physicalprocessor
	Clockspeed                        string      `json:"clockspeed" gorm:"column:clockspeed"`                                                       // clockspeed
	Memory                            StringFloat `json:"memory" gorm:"column:memory"`                                                               // memory
	Storage                           string      `json:"storage" gorm:"column:storage"`                                                             // storage
	Networkperformance                string      `json:"networkperformance" gorm:"column:networkperformance"`                                       // networkperformance
	Operatingsystem                   Os          `json:"operatingsystem" gorm:"column:operatingsystem"`                                             // operatingsystem
	Preinstalledsw                    string      `json:"preinstalledsw" gorm:"column:preinstalledsw"`                                               // preinstalledsw
	Licensemodel                      string      `json:"licensemodel" gorm:"column:licensemodel"`                                                   // licensemodel
	Capacitystatus                    string      `json:"capacitystatus" gorm:"column:capacitystatus"`                                               // capacitystatus
	Tenancy                           string      `json:"tenancy" gorm:"column:tenancy"`                                                             // tenancy
	Dedicatedebsthroughput            string      `json:"dedicatedebsthroughput" gorm:"column:dedicatedebsthroughput"`                               // dedicatedebsthroughput
	Ecu                               StringFloat `json:"ecu" gorm:"column:ecu"`                                                                     // ecu
	Gpumemory                         StringFloat `json:"gpumemory" gorm:"column:gpumemory"`                                                         // gpumemory
	Marketoption                      string      `json:"marketoption" gorm:"column:marketoption"`                                                   // marketoption
	Processorfeatures                 string      `json:"processorfeatures" gorm:"column:processorfeatures"`                                         // processorfeatures
	Ondemandprice                     StringFloat `json:"ondemandprice" gorm:"column:ondemandprice"`                                                 // ondemandprice
	OneYearReservedStandardPrice      StringFloat `json:"one_year_reserved_standard_price" gorm:"column:one_year_reserved_standard_price"`           // one_year_reserved_standard_price
	ThreeYearReservedStandardPrice    StringFloat `json:"three_year_reserved_standard_price" gorm:"column:three_year_reserved_standard_price"`       // three_year_reserved_standard_price
	OneYearReservedConvertiblePrice   StringFloat `json:"one_year_reserved_convertible_price" gorm:"column:one_year_reserved_convertible_price"`     // one_year_reserved_convertible_price
	ThreeYearReservedConvertiblePrice StringFloat `json:"three_year_reserved_convertible_price" gorm:"column:three_year_reserved_convertible_price"` // three_year_reserved_convertible_price
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [AwsEc2Inctance] exists in the database.
func (ae *AwsEc2Inctance) Exists() bool {
	return ae._exists
}

// Deleted returns true when the [AwsEc2Inctance] has been marked for deletion
// from the database.
func (ae *AwsEc2Inctance) Deleted() bool {
	return ae._deleted
}

// Insert inserts the [AwsEc2Inctance] to the database.
func (ae *AwsEc2Inctance) Insert(ctx context.Context, db DB) error {
	switch {
	case ae._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ae._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.aws_ec2_inctances (` +
		`id, regioncode, instancetype, instancefamily, vcpu, physicalprocessor, clockspeed, memory, storage, networkperformance, operatingsystem, preinstalledsw, licensemodel, capacitystatus, tenancy, dedicatedebsthroughput, ecu, gpumemory, marketoption, processorfeatures, ondemandprice, one_year_reserved_standard_price, three_year_reserved_standard_price, one_year_reserved_convertible_price, three_year_reserved_convertible_price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25` +
		`)`
	// run
	logf(sqlstr, ae.ID, ae.Regioncode, ae.Instancetype, ae.Instancefamily, ae.Vcpu, ae.Physicalprocessor, ae.Clockspeed, ae.Memory, ae.Storage, ae.Networkperformance, ae.Operatingsystem, ae.Preinstalledsw, ae.Licensemodel, ae.Capacitystatus, ae.Tenancy, ae.Dedicatedebsthroughput, ae.Ecu, ae.Gpumemory, ae.Marketoption, ae.Processorfeatures, ae.Ondemandprice, ae.OneYearReservedStandardPrice, ae.ThreeYearReservedStandardPrice, ae.OneYearReservedConvertiblePrice, ae.ThreeYearReservedConvertiblePrice)
	if _, err := db.ExecContext(ctx, sqlstr, ae.ID, ae.Regioncode, ae.Instancetype, ae.Instancefamily, ae.Vcpu, ae.Physicalprocessor, ae.Clockspeed, ae.Memory, ae.Storage, ae.Networkperformance, ae.Operatingsystem, ae.Preinstalledsw, ae.Licensemodel, ae.Capacitystatus, ae.Tenancy, ae.Dedicatedebsthroughput, ae.Ecu, ae.Gpumemory, ae.Marketoption, ae.Processorfeatures, ae.Ondemandprice, ae.OneYearReservedStandardPrice, ae.ThreeYearReservedStandardPrice, ae.OneYearReservedConvertiblePrice, ae.ThreeYearReservedConvertiblePrice); err != nil {
		return logerror(err)
	}
	// set exists
	ae._exists = true
	return nil
}

// Update updates a [AwsEc2Inctance] in the database.
func (ae *AwsEc2Inctance) Update(ctx context.Context, db DB) error {
	switch {
	case !ae._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ae._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.aws_ec2_inctances SET ` +
		`regioncode = $1, instancetype = $2, instancefamily = $3, vcpu = $4, physicalprocessor = $5, clockspeed = $6, memory = $7, storage = $8, networkperformance = $9, operatingsystem = $10, preinstalledsw = $11, licensemodel = $12, capacitystatus = $13, tenancy = $14, dedicatedebsthroughput = $15, ecu = $16, gpumemory = $17, marketoption = $18, processorfeatures = $19, ondemandprice = $20, one_year_reserved_standard_price = $21, three_year_reserved_standard_price = $22, one_year_reserved_convertible_price = $23, three_year_reserved_convertible_price = $24 ` +
		`WHERE id = $25`
	// run
	logf(sqlstr, ae.Regioncode, ae.Instancetype, ae.Instancefamily, ae.Vcpu, ae.Physicalprocessor, ae.Clockspeed, ae.Memory, ae.Storage, ae.Networkperformance, ae.Operatingsystem, ae.Preinstalledsw, ae.Licensemodel, ae.Capacitystatus, ae.Tenancy, ae.Dedicatedebsthroughput, ae.Ecu, ae.Gpumemory, ae.Marketoption, ae.Processorfeatures, ae.Ondemandprice, ae.OneYearReservedStandardPrice, ae.ThreeYearReservedStandardPrice, ae.OneYearReservedConvertiblePrice, ae.ThreeYearReservedConvertiblePrice, ae.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ae.Regioncode, ae.Instancetype, ae.Instancefamily, ae.Vcpu, ae.Physicalprocessor, ae.Clockspeed, ae.Memory, ae.Storage, ae.Networkperformance, ae.Operatingsystem, ae.Preinstalledsw, ae.Licensemodel, ae.Capacitystatus, ae.Tenancy, ae.Dedicatedebsthroughput, ae.Ecu, ae.Gpumemory, ae.Marketoption, ae.Processorfeatures, ae.Ondemandprice, ae.OneYearReservedStandardPrice, ae.ThreeYearReservedStandardPrice, ae.OneYearReservedConvertiblePrice, ae.ThreeYearReservedConvertiblePrice, ae.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [AwsEc2Inctance] to the database.
func (ae *AwsEc2Inctance) Save(ctx context.Context, db DB) error {
	if ae.Exists() {
		return ae.Update(ctx, db)
	}
	return ae.Insert(ctx, db)
}

// Upsert performs an upsert for [AwsEc2Inctance].
func (ae *AwsEc2Inctance) Upsert(ctx context.Context, db DB) error {
	switch {
	case ae._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.aws_ec2_inctances (` +
		`id, regioncode, instancetype, instancefamily, vcpu, physicalprocessor, clockspeed, memory, storage, networkperformance, operatingsystem, preinstalledsw, licensemodel, capacitystatus, tenancy, dedicatedebsthroughput, ecu, gpumemory, marketoption, processorfeatures, ondemandprice, one_year_reserved_standard_price, three_year_reserved_standard_price, one_year_reserved_convertible_price, three_year_reserved_convertible_price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`regioncode = EXCLUDED.regioncode, instancetype = EXCLUDED.instancetype, instancefamily = EXCLUDED.instancefamily, vcpu = EXCLUDED.vcpu, physicalprocessor = EXCLUDED.physicalprocessor, clockspeed = EXCLUDED.clockspeed, memory = EXCLUDED.memory, storage = EXCLUDED.storage, networkperformance = EXCLUDED.networkperformance, operatingsystem = EXCLUDED.operatingsystem, preinstalledsw = EXCLUDED.preinstalledsw, licensemodel = EXCLUDED.licensemodel, capacitystatus = EXCLUDED.capacitystatus, tenancy = EXCLUDED.tenancy, dedicatedebsthroughput = EXCLUDED.dedicatedebsthroughput, ecu = EXCLUDED.ecu, gpumemory = EXCLUDED.gpumemory, marketoption = EXCLUDED.marketoption, processorfeatures = EXCLUDED.processorfeatures, ondemandprice = EXCLUDED.ondemandprice, one_year_reserved_standard_price = EXCLUDED.one_year_reserved_standard_price, three_year_reserved_standard_price = EXCLUDED.three_year_reserved_standard_price, one_year_reserved_convertible_price = EXCLUDED.one_year_reserved_convertible_price, three_year_reserved_convertible_price = EXCLUDED.three_year_reserved_convertible_price `
	// run
	logf(sqlstr, ae.ID, ae.Regioncode, ae.Instancetype, ae.Instancefamily, ae.Vcpu, ae.Physicalprocessor, ae.Clockspeed, ae.Memory, ae.Storage, ae.Networkperformance, ae.Operatingsystem, ae.Preinstalledsw, ae.Licensemodel, ae.Capacitystatus, ae.Tenancy, ae.Dedicatedebsthroughput, ae.Ecu, ae.Gpumemory, ae.Marketoption, ae.Processorfeatures, ae.Ondemandprice, ae.OneYearReservedStandardPrice, ae.ThreeYearReservedStandardPrice, ae.OneYearReservedConvertiblePrice, ae.ThreeYearReservedConvertiblePrice)
	if _, err := db.ExecContext(ctx, sqlstr, ae.ID, ae.Regioncode, ae.Instancetype, ae.Instancefamily, ae.Vcpu, ae.Physicalprocessor, ae.Clockspeed, ae.Memory, ae.Storage, ae.Networkperformance, ae.Operatingsystem, ae.Preinstalledsw, ae.Licensemodel, ae.Capacitystatus, ae.Tenancy, ae.Dedicatedebsthroughput, ae.Ecu, ae.Gpumemory, ae.Marketoption, ae.Processorfeatures, ae.Ondemandprice, ae.OneYearReservedStandardPrice, ae.ThreeYearReservedStandardPrice, ae.OneYearReservedConvertiblePrice, ae.ThreeYearReservedConvertiblePrice); err != nil {
		return logerror(err)
	}
	// set exists
	ae._exists = true
	return nil
}

// Delete deletes the [AwsEc2Inctance] from the database.
func (ae *AwsEc2Inctance) Delete(ctx context.Context, db DB) error {
	switch {
	case !ae._exists: // doesn't exist
		return nil
	case ae._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.aws_ec2_inctances ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, ae.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ae.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ae._deleted = true
	return nil
}

// AwsEc2InctanceByID retrieves a row from 'public.aws_ec2_inctances' as a [AwsEc2Inctance].
//
// Generated from index 'aws_ec2_inctances_pkey'.
func AwsEc2InctanceByID(ctx context.Context, db DB, id string) (*AwsEc2Inctance, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, regioncode, instancetype, instancefamily, vcpu, physicalprocessor, clockspeed, memory, storage, networkperformance, operatingsystem, preinstalledsw, licensemodel, capacitystatus, tenancy, dedicatedebsthroughput, ecu, gpumemory, marketoption, processorfeatures, ondemandprice, one_year_reserved_standard_price, three_year_reserved_standard_price, one_year_reserved_convertible_price, three_year_reserved_convertible_price ` +
		`FROM public.aws_ec2_inctances ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	ae := AwsEc2Inctance{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ae.ID, &ae.Regioncode, &ae.Instancetype, &ae.Instancefamily, &ae.Vcpu, &ae.Physicalprocessor, &ae.Clockspeed, &ae.Memory, &ae.Storage, &ae.Networkperformance, &ae.Operatingsystem, &ae.Preinstalledsw, &ae.Licensemodel, &ae.Capacitystatus, &ae.Tenancy, &ae.Dedicatedebsthroughput, &ae.Ecu, &ae.Gpumemory, &ae.Marketoption, &ae.Processorfeatures, &ae.Ondemandprice, &ae.OneYearReservedStandardPrice, &ae.ThreeYearReservedStandardPrice, &ae.OneYearReservedConvertiblePrice, &ae.ThreeYearReservedConvertiblePrice); err != nil {
		return nil, logerror(err)
	}
	return &ae, nil
}
