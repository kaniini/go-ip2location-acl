package ip2location_acl

import (
	"github.com/kaniini/go-confparse"
	"github.com/thekvs/go-net-radix"
)

type IP2LocationDB struct {
	tree *netradix.NetRadixTree
}

func OpenIP2LocationDB(ipv4_db *string, ipv6_db *string) (*IP2LocationDB, error) {
	db := &IP2LocationDB {nil}

	tree, error := netradix.NewNetRadixTree()
	if tree == nil {
		return nil, error
	}

	db.tree = tree

	if ipv4_db != nil {
		db.LoadIP2LocationDB(*ipv4_db)
	}

	if ipv6_db != nil {
		db.LoadIP2LocationDB(*ipv6_db)
	}

	return db, nil
}

func (db *IP2LocationDB) LoadIP2LocationDB (db_path string) {
	cfptr := confparse.LoadConfigFile(db_path)
	if cfptr == nil {
		return
	}

	for ceptr := cfptr.Entries; ceptr != nil; ceptr = ceptr.Next {
		if ceptr.Entries != nil {
			db.ParseACLBlock(ceptr.VarData, ceptr.Entries)
		}
	}
}

func (db *IP2LocationDB) ParseACLBlock (country string, subblock *confparse.ConfigEntry) {
	for ceptr := subblock.Entries; ceptr != nil; ceptr = ceptr.Next {
		if ceptr.VarName[0:7] == "::FFFF:" {
			continue;
		}
		db.tree.Add(ceptr.VarName, country)
	}
}

func (db *IP2LocationDB) Close () {
	db.tree.Close()
}

func (db *IP2LocationDB) Lookup (addr string) (country string) {
	status, country, err := db.tree.SearchBest(addr)
	if err != nil {
		return "???"
	}
	if status {
		return country
	}
	return "???"
}
