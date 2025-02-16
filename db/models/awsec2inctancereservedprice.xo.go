package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AwsEc2InctanceReservedPrice represents a row from 'public.aws_ec2_inctance_reserved_prices'.
type AwsEc2InctanceReservedPrice struct {
	ID                  string          `json:"id" gorm:"column:id"`                                   // id
	Leasecontractlength StringInt       `json:"leasecontractlength" gorm:"column:leasecontractlength"` // leasecontractlength
	Purchaseoption      PurchaseOption  `json:"purchaseoption" gorm:"column:purchaseoption"`           // purchaseoption
	Offeringclass       OfferingClass   `json:"offeringclass" gorm:"column:offeringclass"`             // offeringclass
	ReservedPrice       sql.NullFloat64 `json:"reserved_price" gorm:"column:reserved_price"`           // reserved_price
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [AwsEc2InctanceReservedPrice] exists in the database.
func (aerp *AwsEc2InctanceReservedPrice) Exists() bool {
	return aerp._exists
}

// Deleted returns true when the [AwsEc2InctanceReservedPrice] has been marked for deletion
// from the database.
func (aerp *AwsEc2InctanceReservedPrice) Deleted() bool {
	return aerp._deleted
}

// Insert inserts the [AwsEc2InctanceReservedPrice] to the database.
func (aerp *AwsEc2InctanceReservedPrice) Insert(ctx context.Context, db DB) error {
	switch {
	case aerp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case aerp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.aws_ec2_inctance_reserved_prices (` +
		`id, leasecontractlength, purchaseoption, offeringclass, reserved_price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)`
	// run
	logf(sqlstr, aerp.ID, aerp.Leasecontractlength, aerp.Purchaseoption, aerp.Offeringclass, aerp.ReservedPrice)
	if _, err := db.ExecContext(ctx, sqlstr, aerp.ID, aerp.Leasecontractlength, aerp.Purchaseoption, aerp.Offeringclass, aerp.ReservedPrice); err != nil {
		return logerror(err)
	}
	// set exists
	aerp._exists = true
	return nil
}

// Update updates a [AwsEc2InctanceReservedPrice] in the database.
func (aerp *AwsEc2InctanceReservedPrice) Update(ctx context.Context, db DB) error {
	switch {
	case !aerp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case aerp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.aws_ec2_inctance_reserved_prices SET ` +
		`leasecontractlength = $1, purchaseoption = $2, offeringclass = $3, reserved_price = $4 ` +
		`WHERE id = $5`
	// run
	logf(sqlstr, aerp.Leasecontractlength, aerp.Purchaseoption, aerp.Offeringclass, aerp.ReservedPrice, aerp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, aerp.Leasecontractlength, aerp.Purchaseoption, aerp.Offeringclass, aerp.ReservedPrice, aerp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [AwsEc2InctanceReservedPrice] to the database.
func (aerp *AwsEc2InctanceReservedPrice) Save(ctx context.Context, db DB) error {
	if aerp.Exists() {
		return aerp.Update(ctx, db)
	}
	return aerp.Insert(ctx, db)
}

// Upsert performs an upsert for [AwsEc2InctanceReservedPrice].
func (aerp *AwsEc2InctanceReservedPrice) Upsert(ctx context.Context, db DB) error {
	switch {
	case aerp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.aws_ec2_inctance_reserved_prices (` +
		`id, leasecontractlength, purchaseoption, offeringclass, reserved_price` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`leasecontractlength = EXCLUDED.leasecontractlength, purchaseoption = EXCLUDED.purchaseoption, offeringclass = EXCLUDED.offeringclass, reserved_price = EXCLUDED.reserved_price `
	// run
	logf(sqlstr, aerp.ID, aerp.Leasecontractlength, aerp.Purchaseoption, aerp.Offeringclass, aerp.ReservedPrice)
	if _, err := db.ExecContext(ctx, sqlstr, aerp.ID, aerp.Leasecontractlength, aerp.Purchaseoption, aerp.Offeringclass, aerp.ReservedPrice); err != nil {
		return logerror(err)
	}
	// set exists
	aerp._exists = true
	return nil
}

// Delete deletes the [AwsEc2InctanceReservedPrice] from the database.
func (aerp *AwsEc2InctanceReservedPrice) Delete(ctx context.Context, db DB) error {
	switch {
	case !aerp._exists: // doesn't exist
		return nil
	case aerp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.aws_ec2_inctance_reserved_prices ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, aerp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, aerp.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	aerp._deleted = true
	return nil
}

// AwsEc2InctanceReservedPriceByID retrieves a row from 'public.aws_ec2_inctance_reserved_prices' as a [AwsEc2InctanceReservedPrice].
//
// Generated from index 'aws_ec2_inctance_reserved_prices_pkey'.
func AwsEc2InctanceReservedPriceByID(ctx context.Context, db DB, id string) (*AwsEc2InctanceReservedPrice, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, leasecontractlength, purchaseoption, offeringclass, reserved_price ` +
		`FROM public.aws_ec2_inctance_reserved_prices ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	aerp := AwsEc2InctanceReservedPrice{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&aerp.ID, &aerp.Leasecontractlength, &aerp.Purchaseoption, &aerp.Offeringclass, &aerp.ReservedPrice); err != nil {
		return nil, logerror(err)
	}
	return &aerp, nil
}
