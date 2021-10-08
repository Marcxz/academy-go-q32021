package infraestructure

import (
	"database/sql"
	"fmt"

	"github.com/Marcxz/academy-go-q32021/conf"
	_ "github.com/lib/pq"
)

// routergenerator - interface that makes the contract to generate a route from 2 latlong coordinates, returns a route geojson and an error if exists
type routegenerator interface {
	GenerateRoute(float64, float64, float64, float64) (string, error)
}

// starter - interface that makes the contract to init a db connection
type starter interface {
	InitDB() error
}

// GeoDB - interface that we use to concatenate all infraestructure interfaces
type GeoDB interface {
	starter
	routegenerator
}

// geoDB - the struct that we use to isolate infraestructure with repository
type geoDB struct {
	Db  *sql.DB
	con *conf.Config
}

// NewGeoDB - constructor func to generate a link between infraestructure with repository, returns a GeoDB interface
func NewGeoDB(con *conf.Config) GeoDB {
	return &geoDB{
		Db:  nil,
		con: con,
	}
}

// InitDB - func to open a postgreSQL database connection, retuns an error if exist
func (gdb *geoDB) InitDB() error {
	var err error

	con := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", gdb.con.PostgresUser, gdb.con.PostgresPassword, gdb.con.PostgresHost, gdb.con.PostgresPort, gdb.con.PostgresDb)
	gdb.Db, err = sql.Open("postgres", con)

	if err != nil {
		return err
	}

	err = gdb.Db.Ping()

	if err != nil {
		return err
	}
	return nil
}

// GenerateRoute - Func wich recieve 2 latlng Coordinate, applies Dijstra algorithm, returns a route string and/or error if exist
func (gdb *geoDB) GenerateRoute(latA float64, lngA float64, latB float64, lngB float64) (string, error) {

	err := gdb.InitDB()
	if err != nil {
		return "", err
	}

	q := `
	with ruta as (
	select * from pgr_dijkstra('Select gid as id, source, target, length_m as cost 
		From public."taller_pgRouting_ways" tprw 
	', 
	
	(select id from public."taller_pgRouting_ways_vertices_pgr" tprwvp
	order by st_distance(st_transform(geom, 32613), st_transform(ST_SetSRID(st_makepoint($1, $2), 4326), 32613)) asc
	limit 1), 
	
	(select id from public."taller_pgRouting_ways_vertices_pgr" tprwvp
	order by st_distance(st_transform(geom, 32613), st_transform(ST_SetSRID(st_makepoint($3, $4), 4326), 32613)) asc
	limit 1), 
	
	directed := true))
	select st_asgeojson(tprw.geom) geojson from ruta left join public."taller_pgRouting_ways" tprw on ruta.edge = tprw.gid
	where st_asgeojson(tprw.geom) is not null`

	rows, err := gdb.Db.Query(q, lngA, latA, lngB, latB)
	if err != nil {
		return "", err
	}

	defer rows.Close()
	route := ""
	for rows.Next() {
		r := ""
		err = rows.Scan(&r)
		if err != nil {
			return "", err
		}
		route += fmt.Sprintf("%s,", r)
	}
	if len(route) > 0 {
		route = route[:len(route)-1]
	}
	return route, nil
}
