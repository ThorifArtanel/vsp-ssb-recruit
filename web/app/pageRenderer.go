package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Spesifications struct {
	Id        string
	Name      string
	IsActive  bool
	Infos     Info
	JobDesc   []string
	QGeneric  []string
	QSpecific []string
}

type Info struct {
	Level      string
	Kebutuhan  string
	Penempatan string
}

func PageRenderer(c *gin.Context) {
	spesifications := []Spesifications{
		{
			Id:       "recruitment-selection",
			Name:     "Recruitment & Selection",
			IsActive: true,
			Infos: Info{
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melakukan job posting",
				"Melakukan screening dan seleksi calon karyawan",
				"Mengolah data dan hasil seleksi calon karyawan",
				"Membuat konten flyer (job vacancy)",
			},
			QGeneric: []string{
				"Menguasai Ms Office",
				"Mampu berkomunikasi dengan baik",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Kreatif, ramah, sopan dan inisiatif",
				"Lebih disukai yang berpengalaman",
			},
			QSpecific: []string{
				"Menguasai Ms Office",
				"Mampu berkomunikasi dengan baik",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Kreatif, ramah, sopan dan inisiatif",
				"Lebih disukai yang berpengalaman",
			},
		},

		{
			Id:       "assessment",
			Name:     "Assessment",
			IsActive: false,
			Infos: Info{
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melakukan pemetaan potensi karyawan",
				"Melaksanakan proses assessment untuk promosi dan mutasi",
				"Menyusun laporan hasil assessment untuk stakeholder",
			},
			QGeneric: []string{
				"Usia maksimal 35 tahun",
				"IPK minimal 3.0 (swasta) atau 2.75 (negeri)",
				"Penempatan: Solo",
				"Mampu bekerja dengan team maupun individu",
			},
			QSpecific: []string{
				"Mampu menggunakan tools assessment",
				"S1 Psikologi",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Lebih disukai yang berpengalaman",
			},
		},

		{
			Id:       "ssb-corp-university",
			Name:     "SSB Corp University",
			IsActive: false,
			Infos: Info{
				Level:      "Staff",
				Kebutuhan:  "2",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menyusun surat menyurat",
				"Melaksanakan pelatihan internal perusahaan",
				"Menjalankan training booth camp",
				"Menyusun laporan pelaksanaan training",
				"Menjalankan fungsi event organizer dalam acara internal perusahaan",
			},
			QGeneric: []string{
				"Usia maksimal 35 tahun",
				"Pendidikan minimal S1 semua jurusan",
				"IPK minimal 3.0 (swasta) atau 2.75 (negeri)",
				"Penempatan: Solo",
				"Mampu bekerja dengan team maupun individu",
			},
			QSpecific: []string{
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Menguasai Ms Office",
				"Mampu berkomunikasi didepan forum",
			},
		},
	}
	c.HTML(http.StatusOK, "index.html", spesifications)
}
