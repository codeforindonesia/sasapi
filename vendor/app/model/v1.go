package model

import (
	"app/shared/database"
	"log"
)

type Realisasi struct {
	Kdoutput string         `db:"kdoutput"  json:"kdoutput"`
	KodeKomponen		*string 		  `db:"kdkmpnen"  json:"kdkmpnen"`
	KodeSubKomponen		*string 		  `db:"kdskmpnen"  json:"kdskmpnen"`
	Kdakun   string         `db:"kdakun"  json:"kdakun"`
	Nosp2d   string         `db:"nosp2d"  json:"nosp2d"`
	Tgsp2d   string         `db:"tgsp2d"  json:"tgsp2d"`
	Nospp    string         `db:"nospp"  json:"nospp"`
	Tgspp string `db:"tgspp"  json:"tgspp"`
	Nospm string `db:"nospm"  json:"nospm"`
	Tgspm    string `db:"tgspm"  json:"tgspm"`
	Kdnpwp   string `db:"kdnpwp2"  json:"npwp"`
	Penyedia string `db:"uraiben1"  json:"penyedia"`
	Alamat   string `db:"uraiben2"  json:"alamat"`
	Nominal		string 		  `db:"nilmak"  json:"nominal"`
}
type Pagu struct {
	KodeGiat		string 		  `db:"kdgiat"  json:"kdgiat"`
	KodeOutput		string 		  `db:"kdoutput"  json:"kdoutput"`
	KodeSubOutput		string 		  `db:"kdsoutput"  json:"kdsoutput"`
	KodeKomponen		string 		  `db:"kdkmpnen"  json:"kdkmpnen"`
	KodeSubKomponen		string 		  `db:"kdskmpnen"  json:"kdskmpnen"`
	KodeAkun		string 		  `db:"kdakun"  json:"kdakun"`
	NamaItem		string 		  `db:"nmitem"  json:"nmitem"`
	Jumlah		string 		  `db:"jumlah"  json:"jumlah"`
}

func RealisasiByDate(startDate, endDate string) ([]Realisasi, error) {
	var err error
	var result []Realisasi
	err = database.SQL.Select(&result, "SELECT m_spmmak.kdoutput as kdoutput, d_item.kdkmpnen as kdkmpnen, d_item.kdskmpnen as kdskmpnen, m_spmmak.kdakun as kdakun, m_spmmak.nosp2d as nosp2d, m_spmmak.tgsp2d as tgsp2d, CONCAT(m_spmind.nospp, m_spmind.nospp2) AS nospp, m_spmind.tgspp as tgspp, CONCAT(m_spmind.nospm, m_spmind.nospm2) AS nospm, m_spmind.tgspm as tgspm, m_spminfo.kdnpwp2 as kdnpwp2, m_spminfo.uraiben1 as uraiben1, m_spminfo.uraiben2 as uraiben2, m_spmmak.nilmak as nilmak FROM m_spmind LEFT JOIN m_spmmak ON m_spmind.nosp2d = m_spmmak.nosp2d LEFT JOIN m_spminfo ON m_spmind.nosp2d = m_spminfo.nosp2d LEFT JOIN d_item ON d_item.kdoutput = m_spmmak.kdoutput AND d_item.kdbeban = m_spmmak.kdbeban AND d_item.kdakun = m_spmmak.kdakun AND d_item.kdgiat = m_spmmak.kdgiat WHERE m_spmind.tgspm >= ? AND m_spmind.tgspm <= ? GROUP BY m_spmmak.nosp2d", startDate, endDate)
	if err != nil {
		log.Println(err)
	}
	return result, standardizeError(err)
}

func PaguByDate() ([]Pagu, error) {
	var err error
	var result []Pagu
	err = database.SQL.Select(&result, "SELECT kdgiat,  kdoutput,  kdsoutput,  kdkmpnen,  kdskmpnen,  kdakun,  nmitem, jumlah FROM d_item ORDER BY kdgiat, kdoutput, kdsoutput, kdkmpnen, kdskmpnen, kdakun")
	if err != nil {
		log.Println(err)
	}
	return result, standardizeError(err)
}


