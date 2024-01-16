package postgres

import (
	"context"
	"em-test-task/internal/domain/entity"
	"em-test-task/internal/domain/storages"
	"em-test-task/internal/domain/storages/dto"
	"em-test-task/pkg/db/postgresql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type BioInfoRepository storages.BioInfoStorage

type bioInfoRepository struct {
	pool postgresql.PGXPool

	log *logrus.Entry
}

func NewBioInfoRepository(pool postgresql.PGXPool, log *logrus.Entry) BioInfoRepository {
	return &bioInfoRepository{pool: pool, log: log}
}

func (bIR *bioInfoRepository) GetByFilters(ctx context.Context, filters map[string]string) ([]*entity.BioInfo, error) {
	b := sq.Select("id", "name", "surname", "patronymic", "age", "gender", "national").PlaceholderFormat(sq.Dollar)
	for name, value := range filters {
		b = b.Where(sq.Eq{name: value})
	}

	query, _, err := b.ToSql()
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}

	rows, err := bIR.pool.Query(
		ctx,
		query,
	)
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}
	defer rows.Close()

	bIs := make([]*entity.BioInfo, 0)
	err = pgxscan.ScanAll(bIs, rows)
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}

	return bIs, nil
}

func (bIR *bioInfoRepository) Store(ctx context.Context, bIC *dto.BioInfoCreate) (*entity.BioInfo, error) {
	var id int
	err := bIR.pool.QueryRow(
		ctx,
		"INSERT INTO bioinfo(name, surname, patronymic, age, gender, national) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		bIC.Name,
		bIC.Surname,
		bIC.Patronymic,
		bIC.Age,
		bIC.Gender,
		bIC.National,
	).Scan(&id)
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}

	return bIC.ToEntity(id), nil
}

func (bIR *bioInfoRepository) GetByID(ctx context.Context, id int) (*entity.BioInfo, error) {
	var bI *entity.BioInfo
	err := bIR.pool.QueryRow(
		ctx,
		"SELECT id, name, surname, patronymic, age, gender, national from bioinfo where id = $1",
		id,
	).Scan(
		bI.ID,
		bI.Name,
		bI.Surname,
		bI.Patronymic,
		bI.Age,
		bI.Gender,
		bI.National,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}

	return bI, nil
}

func (bIR *bioInfoRepository) GetAll(ctx context.Context) ([]*entity.BioInfo, error) {
	rows, err := bIR.pool.Query(
		ctx,
		"SELECT id, name, surname, patronymic, age, gender, national from bioinfo",
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return []*entity.BioInfo{}, nil
	}
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}
	defer rows.Close()

	bIs := make([]*entity.BioInfo, 0)

	err = pgxscan.ScanAll(&bIs, rows)
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}

	return bIs, nil
}

func (bIR *bioInfoRepository) UpdateByID(ctx context.Context, bI *entity.BioInfo) error {
	_, err := bIR.pool.Exec(
		ctx,
		"UPDATE bioinfo set name=$2, surname = $3, patronymic = $4, age = $5, gender = $6, national = $7 where id = $1",
		bI.ID,
		bI.Name,
		bI.Surname,
		bI.Patronymic,
		bI.Age,
		bI.Gender,
		bI.National,
	)
	if err != nil {
		bIR.log.Errorln(err)
		return err
	}

	return nil
}

func (bIR *bioInfoRepository) DeleteByID(ctx context.Context, id int) (*entity.BioInfo, error) {
	var bI entity.BioInfo
	err := bIR.pool.QueryRow(
		ctx,
		"DELETE  FROM bioinfo where id = $1 RETURNING id, name, surname, patronymic, age, gender, national",
		id,
	).Scan(
		&bI.ID,
		&bI.Name,
		&bI.Surname,
		&bI.Patronymic,
		&bI.Age,
		&bI.Gender,
		&bI.National,
	)
	if err != nil {
		bIR.log.Errorln(err)
		return nil, err
	}

	return &bI, nil
}
