// Code generated by pggen. DO NOT EDIT.

package pggen

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const insertJobSQL = `INSERT INTO jobs (
    run_id,
    phase,
    status
) VALUES (
    $1,
    $2,
    $3
);`

type InsertJobParams struct {
	RunID  pgtype.Text
	Phase  pgtype.Text
	Status pgtype.Text
}

// InsertJob implements Querier.InsertJob.
func (q *DBQuerier) InsertJob(ctx context.Context, params InsertJobParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "InsertJob")
	cmdTag, err := q.conn.Exec(ctx, insertJobSQL, params.RunID, params.Phase, params.Status)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query InsertJob: %w", err)
	}
	return cmdTag, err
}

// InsertJobBatch implements Querier.InsertJobBatch.
func (q *DBQuerier) InsertJobBatch(batch genericBatch, params InsertJobParams) {
	batch.Queue(insertJobSQL, params.RunID, params.Phase, params.Status)
}

// InsertJobScan implements Querier.InsertJobScan.
func (q *DBQuerier) InsertJobScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec InsertJobBatch: %w", err)
	}
	return cmdTag, err
}

const findJobsSQL = `SELECT
    j.run_id,
    j.phase,
    j.status,
    j.signaled,
    j.agent_id,
    w.agent_pool_id,
    r.workspace_id,
    w.organization_name
FROM jobs j
JOIN runs r USING (run_id)
JOIN workspaces w USING (workspace_id)
;`

type FindJobsRow struct {
	RunID            pgtype.Text `json:"run_id"`
	Phase            pgtype.Text `json:"phase"`
	Status           pgtype.Text `json:"status"`
	Signaled         pgtype.Bool `json:"signaled"`
	AgentID          pgtype.Text `json:"agent_id"`
	AgentPoolID      pgtype.Text `json:"agent_pool_id"`
	WorkspaceID      pgtype.Text `json:"workspace_id"`
	OrganizationName pgtype.Text `json:"organization_name"`
}

// FindJobs implements Querier.FindJobs.
func (q *DBQuerier) FindJobs(ctx context.Context) ([]FindJobsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindJobs")
	rows, err := q.conn.Query(ctx, findJobsSQL)
	if err != nil {
		return nil, fmt.Errorf("query FindJobs: %w", err)
	}
	defer rows.Close()
	items := []FindJobsRow{}
	for rows.Next() {
		var item FindJobsRow
		if err := rows.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
			return nil, fmt.Errorf("scan FindJobs row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindJobs rows: %w", err)
	}
	return items, err
}

// FindJobsBatch implements Querier.FindJobsBatch.
func (q *DBQuerier) FindJobsBatch(batch genericBatch) {
	batch.Queue(findJobsSQL)
}

// FindJobsScan implements Querier.FindJobsScan.
func (q *DBQuerier) FindJobsScan(results pgx.BatchResults) ([]FindJobsRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindJobsBatch: %w", err)
	}
	defer rows.Close()
	items := []FindJobsRow{}
	for rows.Next() {
		var item FindJobsRow
		if err := rows.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
			return nil, fmt.Errorf("scan FindJobsBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindJobsBatch rows: %w", err)
	}
	return items, err
}

const findJobSQL = `SELECT
    j.run_id,
    j.phase,
    j.status,
    j.signaled,
    j.agent_id,
    w.agent_pool_id,
    r.workspace_id,
    w.organization_name
FROM jobs j
JOIN runs r USING (run_id)
JOIN workspaces w USING (workspace_id)
WHERE run_id = $1
AND   phase = $2
;`

type FindJobRow struct {
	RunID            pgtype.Text `json:"run_id"`
	Phase            pgtype.Text `json:"phase"`
	Status           pgtype.Text `json:"status"`
	Signaled         pgtype.Bool `json:"signaled"`
	AgentID          pgtype.Text `json:"agent_id"`
	AgentPoolID      pgtype.Text `json:"agent_pool_id"`
	WorkspaceID      pgtype.Text `json:"workspace_id"`
	OrganizationName pgtype.Text `json:"organization_name"`
}

// FindJob implements Querier.FindJob.
func (q *DBQuerier) FindJob(ctx context.Context, runID pgtype.Text, phase pgtype.Text) (FindJobRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindJob")
	row := q.conn.QueryRow(ctx, findJobSQL, runID, phase)
	var item FindJobRow
	if err := row.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
		return item, fmt.Errorf("query FindJob: %w", err)
	}
	return item, nil
}

// FindJobBatch implements Querier.FindJobBatch.
func (q *DBQuerier) FindJobBatch(batch genericBatch, runID pgtype.Text, phase pgtype.Text) {
	batch.Queue(findJobSQL, runID, phase)
}

// FindJobScan implements Querier.FindJobScan.
func (q *DBQuerier) FindJobScan(results pgx.BatchResults) (FindJobRow, error) {
	row := results.QueryRow()
	var item FindJobRow
	if err := row.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
		return item, fmt.Errorf("scan FindJobBatch row: %w", err)
	}
	return item, nil
}

const findJobForUpdateSQL = `SELECT
    j.run_id,
    j.phase,
    j.status,
    j.signaled,
    j.agent_id,
    w.agent_pool_id,
    r.workspace_id,
    w.organization_name
FROM jobs j
JOIN runs r USING (run_id)
JOIN workspaces w USING (workspace_id)
WHERE run_id = $1
AND   phase = $2
FOR UPDATE OF j
;`

type FindJobForUpdateRow struct {
	RunID            pgtype.Text `json:"run_id"`
	Phase            pgtype.Text `json:"phase"`
	Status           pgtype.Text `json:"status"`
	Signaled         pgtype.Bool `json:"signaled"`
	AgentID          pgtype.Text `json:"agent_id"`
	AgentPoolID      pgtype.Text `json:"agent_pool_id"`
	WorkspaceID      pgtype.Text `json:"workspace_id"`
	OrganizationName pgtype.Text `json:"organization_name"`
}

// FindJobForUpdate implements Querier.FindJobForUpdate.
func (q *DBQuerier) FindJobForUpdate(ctx context.Context, runID pgtype.Text, phase pgtype.Text) (FindJobForUpdateRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindJobForUpdate")
	row := q.conn.QueryRow(ctx, findJobForUpdateSQL, runID, phase)
	var item FindJobForUpdateRow
	if err := row.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
		return item, fmt.Errorf("query FindJobForUpdate: %w", err)
	}
	return item, nil
}

// FindJobForUpdateBatch implements Querier.FindJobForUpdateBatch.
func (q *DBQuerier) FindJobForUpdateBatch(batch genericBatch, runID pgtype.Text, phase pgtype.Text) {
	batch.Queue(findJobForUpdateSQL, runID, phase)
}

// FindJobForUpdateScan implements Querier.FindJobForUpdateScan.
func (q *DBQuerier) FindJobForUpdateScan(results pgx.BatchResults) (FindJobForUpdateRow, error) {
	row := results.QueryRow()
	var item FindJobForUpdateRow
	if err := row.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
		return item, fmt.Errorf("scan FindJobForUpdateBatch row: %w", err)
	}
	return item, nil
}

const findAllocatedJobsSQL = `SELECT
    j.run_id,
    j.phase,
    j.status,
    j.signaled,
    j.agent_id,
    w.agent_pool_id,
    r.workspace_id,
    w.organization_name
FROM jobs j
JOIN runs r USING (run_id)
JOIN workspaces w USING (workspace_id)
WHERE j.agent_id = $1
AND   j.status = 'allocated';`

type FindAllocatedJobsRow struct {
	RunID            pgtype.Text `json:"run_id"`
	Phase            pgtype.Text `json:"phase"`
	Status           pgtype.Text `json:"status"`
	Signaled         pgtype.Bool `json:"signaled"`
	AgentID          pgtype.Text `json:"agent_id"`
	AgentPoolID      pgtype.Text `json:"agent_pool_id"`
	WorkspaceID      pgtype.Text `json:"workspace_id"`
	OrganizationName pgtype.Text `json:"organization_name"`
}

// FindAllocatedJobs implements Querier.FindAllocatedJobs.
func (q *DBQuerier) FindAllocatedJobs(ctx context.Context, agentID pgtype.Text) ([]FindAllocatedJobsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindAllocatedJobs")
	rows, err := q.conn.Query(ctx, findAllocatedJobsSQL, agentID)
	if err != nil {
		return nil, fmt.Errorf("query FindAllocatedJobs: %w", err)
	}
	defer rows.Close()
	items := []FindAllocatedJobsRow{}
	for rows.Next() {
		var item FindAllocatedJobsRow
		if err := rows.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
			return nil, fmt.Errorf("scan FindAllocatedJobs row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAllocatedJobs rows: %w", err)
	}
	return items, err
}

// FindAllocatedJobsBatch implements Querier.FindAllocatedJobsBatch.
func (q *DBQuerier) FindAllocatedJobsBatch(batch genericBatch, agentID pgtype.Text) {
	batch.Queue(findAllocatedJobsSQL, agentID)
}

// FindAllocatedJobsScan implements Querier.FindAllocatedJobsScan.
func (q *DBQuerier) FindAllocatedJobsScan(results pgx.BatchResults) ([]FindAllocatedJobsRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindAllocatedJobsBatch: %w", err)
	}
	defer rows.Close()
	items := []FindAllocatedJobsRow{}
	for rows.Next() {
		var item FindAllocatedJobsRow
		if err := rows.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
			return nil, fmt.Errorf("scan FindAllocatedJobsBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAllocatedJobsBatch rows: %w", err)
	}
	return items, err
}

const findAndUpdateSignaledJobsSQL = `UPDATE jobs AS j
SET signaled = NULL
FROM runs r, workspaces w
WHERE j.run_id = r.run_id
AND   r.workspace_id = w.workspace_id
AND   j.agent_id = $1
AND   j.status = 'running'
AND   j.signaled IS NOT NULL
RETURNING
    j.run_id,
    j.phase,
    j.status,
    j.signaled,
    j.agent_id,
    w.agent_pool_id,
    r.workspace_id,
    w.organization_name
;`

type FindAndUpdateSignaledJobsRow struct {
	RunID            pgtype.Text `json:"run_id"`
	Phase            pgtype.Text `json:"phase"`
	Status           pgtype.Text `json:"status"`
	Signaled         pgtype.Bool `json:"signaled"`
	AgentID          pgtype.Text `json:"agent_id"`
	AgentPoolID      pgtype.Text `json:"agent_pool_id"`
	WorkspaceID      pgtype.Text `json:"workspace_id"`
	OrganizationName pgtype.Text `json:"organization_name"`
}

// FindAndUpdateSignaledJobs implements Querier.FindAndUpdateSignaledJobs.
func (q *DBQuerier) FindAndUpdateSignaledJobs(ctx context.Context, agentID pgtype.Text) ([]FindAndUpdateSignaledJobsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindAndUpdateSignaledJobs")
	rows, err := q.conn.Query(ctx, findAndUpdateSignaledJobsSQL, agentID)
	if err != nil {
		return nil, fmt.Errorf("query FindAndUpdateSignaledJobs: %w", err)
	}
	defer rows.Close()
	items := []FindAndUpdateSignaledJobsRow{}
	for rows.Next() {
		var item FindAndUpdateSignaledJobsRow
		if err := rows.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
			return nil, fmt.Errorf("scan FindAndUpdateSignaledJobs row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAndUpdateSignaledJobs rows: %w", err)
	}
	return items, err
}

// FindAndUpdateSignaledJobsBatch implements Querier.FindAndUpdateSignaledJobsBatch.
func (q *DBQuerier) FindAndUpdateSignaledJobsBatch(batch genericBatch, agentID pgtype.Text) {
	batch.Queue(findAndUpdateSignaledJobsSQL, agentID)
}

// FindAndUpdateSignaledJobsScan implements Querier.FindAndUpdateSignaledJobsScan.
func (q *DBQuerier) FindAndUpdateSignaledJobsScan(results pgx.BatchResults) ([]FindAndUpdateSignaledJobsRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindAndUpdateSignaledJobsBatch: %w", err)
	}
	defer rows.Close()
	items := []FindAndUpdateSignaledJobsRow{}
	for rows.Next() {
		var item FindAndUpdateSignaledJobsRow
		if err := rows.Scan(&item.RunID, &item.Phase, &item.Status, &item.Signaled, &item.AgentID, &item.AgentPoolID, &item.WorkspaceID, &item.OrganizationName); err != nil {
			return nil, fmt.Errorf("scan FindAndUpdateSignaledJobsBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAndUpdateSignaledJobsBatch rows: %w", err)
	}
	return items, err
}

const updateJobSQL = `UPDATE jobs
SET status   = $1,
    signaled = $2,
    agent_id = $3
WHERE run_id = $4
AND   phase = $5
RETURNING *;`

type UpdateJobParams struct {
	Status   pgtype.Text
	Signaled pgtype.Bool
	AgentID  pgtype.Text
	RunID    pgtype.Text
	Phase    pgtype.Text
}

type UpdateJobRow struct {
	RunID    pgtype.Text `json:"run_id"`
	Phase    pgtype.Text `json:"phase"`
	Status   pgtype.Text `json:"status"`
	AgentID  pgtype.Text `json:"agent_id"`
	Signaled pgtype.Bool `json:"signaled"`
}

// UpdateJob implements Querier.UpdateJob.
func (q *DBQuerier) UpdateJob(ctx context.Context, params UpdateJobParams) (UpdateJobRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateJob")
	row := q.conn.QueryRow(ctx, updateJobSQL, params.Status, params.Signaled, params.AgentID, params.RunID, params.Phase)
	var item UpdateJobRow
	if err := row.Scan(&item.RunID, &item.Phase, &item.Status, &item.AgentID, &item.Signaled); err != nil {
		return item, fmt.Errorf("query UpdateJob: %w", err)
	}
	return item, nil
}

// UpdateJobBatch implements Querier.UpdateJobBatch.
func (q *DBQuerier) UpdateJobBatch(batch genericBatch, params UpdateJobParams) {
	batch.Queue(updateJobSQL, params.Status, params.Signaled, params.AgentID, params.RunID, params.Phase)
}

// UpdateJobScan implements Querier.UpdateJobScan.
func (q *DBQuerier) UpdateJobScan(results pgx.BatchResults) (UpdateJobRow, error) {
	row := results.QueryRow()
	var item UpdateJobRow
	if err := row.Scan(&item.RunID, &item.Phase, &item.Status, &item.AgentID, &item.Signaled); err != nil {
		return item, fmt.Errorf("scan UpdateJobBatch row: %w", err)
	}
	return item, nil
}