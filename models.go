package main

// Variable Types Declaration

// LAS_T_SCORING
type Las_t_scoring struct {
    ID_Scoring uint64 `json:"id_scoring"`
    Net_income uint64 `json:"net_income"`
    Angsuran uint64 `json:"angsuran"`
    Nilai_likuidasi_agunan uint64 `json:"nilai_likuidasi_agunan"`
    Plafon uint64 `json:"plafon"`
    Jenis_bukti_kepemilikan_agunan string `json:"jenis_bukti_kepemilikan_agunan"`
    Usia uint `json:"usia"`
    Debitur_baru string `json:"debitur_baru"`
    Lama_usaha uint `json:"lama_usaha"`
    Status_perkawinan string `json:"status_perkawinan"`
    Punya_usaha_sampingan string `json:"punya_usaha_sampingan"`
    Tujuan_penggunaan_kredit string `json:"tujuan_penggunaan_kredit"`
    Punya_pelanggan_tetap string `json:"punya_pelanggan_tetap"`
}

func (l Las_t_scoring) ToClean() Las_t_scoring_clean {
    var l1 Las_t_scoring_clean
    l1.ID_Scoring              = l.ID_Scoring
    l1.Net_income              = l.Net_income
    l1.Angsuran                = l.Angsuran
    l1.Nilai_likuidasi_agunan  = l.Nilai_likuidasi_agunan
    l1.Plafon                  = l.Plafon
    l1.Usia                    = l.Usia
    l1.Lama_usaha              = l.Lama_usaha
    switch {
    case l.Debitur_baru == "Tidak":
        l1.Debitur_lama = 1
    default:
        l1.Debitur_lama = 0
    }

    switch {
    case l.Punya_usaha_sampingan == "Ya":
        l1.Punya_usaha_sampingan_Ya  = 1
    default:
        l1.Punya_usaha_sampingan_Ya  = 0
    }

    switch {
    case l.Punya_pelanggan_tetap == "Ya":
        l1.Punya_pelanggan_tetap_Ya  = 1
    default:
        l1.Punya_pelanggan_tetap_Ya  = 0
    }

    return l1
}

// LAS_T_SCORING Array
type Las_t_scoring_array []Las_t_scoring

func (l Las_t_scoring_array) Get(id uint64) (Las_t_scoring, bool) {
    for _, item := range l {
        if item.ID_Scoring == id {
            return item, true
        }
    }
    return Las_t_scoring{}, false
}

func (l *Las_t_scoring_array) Add(l_new Las_t_scoring) {
    *l = append(*l, l_new)
}

func (l *Las_t_scoring_array) Delete(id uint64) {
    var l1 Las_t_scoring_array
    l1 = *l
    for index, item := range *l {
        if item.ID_Scoring == id {
            l1 = append(l1[:index], l1[index+1:]...)
            *l = l1
            break
        }
    }
}

// LAS_T_SCORING_CLEAN
type Las_t_scoring_clean struct {
    ID_Scoring uint64 `json:"id_scoring"`
    Net_income uint64 `json:"net_income"`
    Angsuran uint64 `json:"angsuran"`
    Nilai_likuidasi_agunan uint64 `json:"nilai_likuidasi_agunan"`
    Plafon uint64 `json:"plafon"`
    // Jenis_bukti_kepemilikan_agunan_Kwitansi uint `json:"jenis_bukti_kepemilikan_agunan_kwitansi"`
    // Jenis_bukti_kepemilikan_agunan_DiluarSertifikat uint `json:"jenis_bukti_kepemilikan_agunan_diluarsertifikat"`
    // Jenis_bukti_kepemilikan_agunan_BPKB uint `json:"jenis_bukti_kepemilikan_agunan_bpkb"`
    // Jenis_bukti_kepemilikan_agunan_Sertifikat uint `json:"jenis_bukti_kepemilikan_agunan_sertifikat"`
    Usia uint `json:"usia"`
    Debitur_lama uint `json:"debitur_lama"`
    Lama_usaha uint `json:"lama_usaha"`
    // Status_perkawinan_Menikah uint `json:"status_perkawinan"`
    Punya_usaha_sampingan_Ya uint `json:"punya_usaha_sampingan"`
    // Tujuan_penggunaan_kredit_ModalKerja uint `json:"tujuan_penggunaan_kredit_modalkerja"`
    // Tujuan_penggunaan_kredit_PenggantiModalKerja uint `json:"tujuan_penggunaan_kredit_penggantimodalkerja"`
    // Tujuan_penggunaan_kredit_Investasi uint `json:"tujuan_penggunaan_kredit_investasi"`
    Punya_pelanggan_tetap_Ya uint `json:"punya_pelanggan_tetap"`
}

func (l Las_t_scoring_clean) Score() Las_status {
    var l1 Las_status
    var score int
    score = 0
    l1.ID_Scoring   = l.ID_Scoring

    ratio_income_to_angsuran := float32(l.Net_income) / float32(l.Angsuran)
    switch {
    case ratio_income_to_angsuran > 3:
        score += 325
    case ratio_income_to_angsuran >= 2:
        score += 290
    case ratio_income_to_angsuran >= 1.33:
        score += 225
    }

    ratio_agunan_to_plafon := float32(l.Nilai_likuidasi_agunan) / float32(l.Plafon)
    switch {
    case ratio_agunan_to_plafon > 2:
        score += 155
    case ratio_agunan_to_plafon > 1.5:
        score += 125
    case ratio_agunan_to_plafon >= 1:
        score += 65
    }

    switch {
    case l.Usia > 50:
        score += 40
    case l.Usia > 45:
        score += 30
    case l.Usia >= 28:
        score += 5
    }

    switch {
    case l.Lama_usaha > 8:
        score += 80
    case l.Lama_usaha > 5:
        score += 70
    case l.Lama_usaha >= 2:
        score += 45
    }

    if l.Debitur_lama > 0 {
        score += 20
    }

    if l.Punya_usaha_sampingan_Ya > 0 {
        score += 45
    }

    if l.Punya_pelanggan_tetap_Ya > 0 {
        score += 60
    }

    l1.Score = float32(score)
    return l1
}

// LAS_T_SCORING_CLEAN Array
type Las_t_scoring_clean_array []Las_t_scoring_clean

func (l Las_t_scoring_clean_array) Get(id uint64) (Las_t_scoring_clean, bool) {
    for _, item := range l {
        if item.ID_Scoring == id {
            return item, true
        }
    }
    return Las_t_scoring_clean{}, false
}

func (l *Las_t_scoring_clean_array) Add(l_new Las_t_scoring_clean) {
    *l = append(*l, l_new)
}

func (l *Las_t_scoring_clean_array) Delete(id uint64) {
    var l1 Las_t_scoring_clean_array
    l1 = *l
    for index, item := range *l {
        if item.ID_Scoring == id {
            l1 = append(l1[:index], l1[index+1:]...)
            *l = l1
            break
        }
    }
}

// LAS_STATUS
type Las_status struct {
    ID_Scoring uint64 `json:"id_scoring" default:nil`
    Score float32 `json:"score" default:nil`
}

// LAS_STATUS Array
type Las_status_array []Las_status

func (l Las_status_array) Get(id uint64) (Las_status, bool) {
    for _, item := range l {
        if item.ID_Scoring == id {
            return item, true
        }
    }
    return Las_status{}, false
}

func (l *Las_status_array) Add(l_new Las_status) {
    *l = append(*l, l_new)
}

func (l *Las_status_array) Delete(id uint64) {
    var l1 Las_status_array
    l1 = *l
    for index, item := range *l {
        if item.ID_Scoring == id {
            l1 = append(l1[:index], l1[index+1:]...)
            *l = l1
            break
        }
    }
}

// Variables Declaration
var las_status Las_status_array

// Functions Declaration
func InitLasStatus() []Las_status {
    var las_status_out []Las_status
	las_status_out = append(las_status_out, Las_status{ID_Scoring: 0, Score: 0})
	return las_status_out
}

// Class Interface Methods
type Las_t interface {
    Get(id uint64) float64
    Delete(id uint64)
}

