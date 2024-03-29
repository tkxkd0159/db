// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Jet is an object representing the database table.
type Jet struct {
	ID      int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	PilotID int    `boil:"pilot_id" json:"pilot_id" toml:"pilot_id" yaml:"pilot_id"`
	Age     int    `boil:"age" json:"age" toml:"age" yaml:"age"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Color   string `boil:"color" json:"color" toml:"color" yaml:"color"`

	R *jetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JetColumns = struct {
	ID      string
	PilotID string
	Age     string
	Name    string
	Color   string
}{
	ID:      "id",
	PilotID: "pilot_id",
	Age:     "age",
	Name:    "name",
	Color:   "color",
}

var JetTableColumns = struct {
	ID      string
	PilotID string
	Age     string
	Name    string
	Color   string
}{
	ID:      "jets.id",
	PilotID: "jets.pilot_id",
	Age:     "jets.age",
	Name:    "jets.name",
	Color:   "jets.color",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var JetWhere = struct {
	ID      whereHelperint
	PilotID whereHelperint
	Age     whereHelperint
	Name    whereHelperstring
	Color   whereHelperstring
}{
	ID:      whereHelperint{field: "\"jets\".\"id\""},
	PilotID: whereHelperint{field: "\"jets\".\"pilot_id\""},
	Age:     whereHelperint{field: "\"jets\".\"age\""},
	Name:    whereHelperstring{field: "\"jets\".\"name\""},
	Color:   whereHelperstring{field: "\"jets\".\"color\""},
}

// JetRels is where relationship names are stored.
var JetRels = struct {
	Pilot string
}{
	Pilot: "Pilot",
}

// jetR is where relationships are stored.
type jetR struct {
	Pilot *Pilot `boil:"Pilot" json:"Pilot" toml:"Pilot" yaml:"Pilot"`
}

// NewStruct creates a new relationship struct
func (*jetR) NewStruct() *jetR {
	return &jetR{}
}

func (r *jetR) GetPilot() *Pilot {
	if r == nil {
		return nil
	}
	return r.Pilot
}

// jetL is where Load methods for each relationship are stored.
type jetL struct{}

var (
	jetAllColumns            = []string{"id", "pilot_id", "age", "name", "color"}
	jetColumnsWithoutDefault = []string{"id", "pilot_id", "age", "name", "color"}
	jetColumnsWithDefault    = []string{}
	jetPrimaryKeyColumns     = []string{"id"}
	jetGeneratedColumns      = []string{}
)

type (
	// JetSlice is an alias for a slice of pointers to Jet.
	// This should almost always be used instead of []Jet.
	JetSlice []*Jet

	jetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jetType                 = reflect.TypeOf(&Jet{})
	jetMapping              = queries.MakeStructMapping(jetType)
	jetPrimaryKeyMapping, _ = queries.BindMapping(jetType, jetMapping, jetPrimaryKeyColumns)
	jetInsertCacheMut       sync.RWMutex
	jetInsertCache          = make(map[string]insertCache)
	jetUpdateCacheMut       sync.RWMutex
	jetUpdateCache          = make(map[string]updateCache)
	jetUpsertCacheMut       sync.RWMutex
	jetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single jet record from the query.
func (q jetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Jet, error) {
	o := &Jet{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for jets")
	}

	return o, nil
}

// All returns all Jet records from the query.
func (q jetQuery) All(ctx context.Context, exec boil.ContextExecutor) (JetSlice, error) {
	var o []*Jet

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Jet slice")
	}

	return o, nil
}

// Count returns the count of all Jet records in the query.
func (q jetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count jets rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q jetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if jets exists")
	}

	return count > 0, nil
}

// Pilot pointed to by the foreign key.
func (o *Jet) Pilot(mods ...qm.QueryMod) pilotQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PilotID),
	}

	queryMods = append(queryMods, mods...)

	return Pilots(queryMods...)
}

// LoadPilot allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (jetL) LoadPilot(ctx context.Context, e boil.ContextExecutor, singular bool, maybeJet interface{}, mods queries.Applicator) error {
	var slice []*Jet
	var object *Jet

	if singular {
		var ok bool
		object, ok = maybeJet.(*Jet)
		if !ok {
			object = new(Jet)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeJet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeJet))
			}
		}
	} else {
		s, ok := maybeJet.(*[]*Jet)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeJet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeJet))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &jetR{}
		}
		args = append(args, object.PilotID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &jetR{}
			}

			for _, a := range args {
				if a == obj.PilotID {
					continue Outer
				}
			}

			args = append(args, obj.PilotID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`pilots`),
		qm.WhereIn(`pilots.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pilot")
	}

	var resultSlice []*Pilot
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pilot")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for pilots")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for pilots")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Pilot = foreign
		if foreign.R == nil {
			foreign.R = &pilotR{}
		}
		foreign.R.Jets = append(foreign.R.Jets, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PilotID == foreign.ID {
				local.R.Pilot = foreign
				if foreign.R == nil {
					foreign.R = &pilotR{}
				}
				foreign.R.Jets = append(foreign.R.Jets, local)
				break
			}
		}
	}

	return nil
}

// SetPilot of the jet to the related item.
// Sets o.R.Pilot to related.
// Adds o to related.R.Jets.
func (o *Jet) SetPilot(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Pilot) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"jets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pilot_id"}),
		strmangle.WhereClause("\"", "\"", 2, jetPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PilotID = related.ID
	if o.R == nil {
		o.R = &jetR{
			Pilot: related,
		}
	} else {
		o.R.Pilot = related
	}

	if related.R == nil {
		related.R = &pilotR{
			Jets: JetSlice{o},
		}
	} else {
		related.R.Jets = append(related.R.Jets, o)
	}

	return nil
}

// Jets retrieves all the records using an executor.
func Jets(mods ...qm.QueryMod) jetQuery {
	mods = append(mods, qm.From("\"jets\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"jets\".*"})
	}

	return jetQuery{q}
}

// FindJet retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJet(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Jet, error) {
	jetObj := &Jet{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"jets\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, jetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from jets")
	}

	return jetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Jet) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no jets provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(jetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	jetInsertCacheMut.RLock()
	cache, cached := jetInsertCache[key]
	jetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			jetAllColumns,
			jetColumnsWithDefault,
			jetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(jetType, jetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jetType, jetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"jets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"jets\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into jets")
	}

	if !cached {
		jetInsertCacheMut.Lock()
		jetInsertCache[key] = cache
		jetInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Jet.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Jet) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	jetUpdateCacheMut.RLock()
	cache, cached := jetUpdateCache[key]
	jetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			jetAllColumns,
			jetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update jets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"jets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jetType, jetMapping, append(wl, jetPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update jets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for jets")
	}

	if !cached {
		jetUpdateCacheMut.Lock()
		jetUpdateCache[key] = cache
		jetUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q jetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for jets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for jets")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"jets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, jetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in jet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all jet")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Jet) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no jets provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(jetColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	jetUpsertCacheMut.RLock()
	cache, cached := jetUpsertCache[key]
	jetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			jetAllColumns,
			jetColumnsWithDefault,
			jetColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			jetAllColumns,
			jetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert jets, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jetPrimaryKeyColumns))
			copy(conflict, jetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"jets\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(jetType, jetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jetType, jetMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert jets")
	}

	if !cached {
		jetUpsertCacheMut.Lock()
		jetUpsertCache[key] = cache
		jetUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Jet record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Jet) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Jet provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jetPrimaryKeyMapping)
	sql := "DELETE FROM \"jets\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from jets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for jets")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q jetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no jetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from jets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for jets")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"jets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from jet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for jets")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Jet) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJet(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"jets\".* FROM \"jets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JetSlice")
	}

	*o = slice

	return nil
}

// JetExists checks if the Jet row exists.
func JetExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"jets\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if jets exists")
	}

	return exists, nil
}
