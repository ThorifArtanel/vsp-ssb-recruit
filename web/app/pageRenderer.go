package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Spesifications struct {
	Id         string
	Name       string
	IsActive   bool
	IsImage    bool
	Infos      Info
	JobDesc    []string
	IsJobDesc  bool
	QGeneric   []string
	QSpecific  []string
	IsSpesific bool
	Image      string
}

type Info struct {
	Divisi     string
	Level      string
	Kebutuhan  string
	Penempatan string
}

func PageRenderer(c *gin.Context) {
	spesifications := []Spesifications{
		{
			Id:       "recruitment-selection-staff-solo",
			Name:     "Recruitment & Selection",
			IsActive: true,
			Infos: Info{
				Divisi:     "Human Capital",
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
			QSpecific: []string{
				"Menguasai Ms Office",
				"Mampu berkomunikasi dengan baik",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Kreatif, ramah, sopan dan inisiatif",
				"Lebih disukai yang berpengalaman",
			},
		},
		{
			Id:   "assessment-staff-solo",
			Name: "Assessment",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melakukan pemetaan potensi karyawan",
				"Melaksanakan proses assessment untuk promosi dan mutasi",
				"Menyusun laporan hasil assessment untuk stakeholder",
			},
			QSpecific: []string{
				"Mampu menggunakan tools assessment",
				"S1 Psikologi",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Lebih disukai yang berpengalaman",
			},
		},
		{
			Id:   "ssb-corp-university-staff-solo",
			Name: "SSB Corp University",
			Infos: Info{
				Divisi:     "Human Capital",
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
			QSpecific: []string{
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Menguasai Ms Office",
				"Mampu berkomunikasi didepan forum",
			},
		},
		{
			Id:   "finance-supervisor-solo",
			Name: "Finance",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menjalankan kebijakan keuangan perusahaan",
				"Menyusun dan mengontrol budget perusahaan",
				"Menyusun jadwal pembayaran dan penagihan hutang piutang perusahaan",
				"Melakukan pencatatan transaksi keuangan, entry jurnal secara akurat dan tepat waktu",
				"Membuat laporan keuangan sesuai ketentuan perusahaan",
				"Menyusun analisa laporan keuangan",
			},
			QSpecific: []string{
				"Minimal D3 Akuntansi/Keuangan/Terkait",
				"Berpengalaman minimal 2 tahun",
				"Memiliki kemampuan analisa dan pengolahan data secara cepat",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
			},
		},
		{
			Id:   "sotfware-engineer-staff-solo",
			Name: "Sotfware Engineer",
			Infos: Info{
				Divisi:     "IT",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Mengelola website perusahaan",
				"Membuat sistem informasi dan analisa data.",
				"Implementasi project perusahaan (Exapro & POS)",
				"Mengontrol pendayagunaan dan pemakaian hardware dan software di perusahaan.",
				"Melakukan instalasi, perawatan dan penyediaan hardware & software",
			},
			QSpecific: []string{
				"S1 Teknik Informatika/Terkait",
				"Siap melakukan perjalanan dinas ke seluruh Indonesia",
				"Menguasai bahasa pemrograman",
				"Menguasai troubleshooting perangkat komputer dan jaringan",
			},
		},
		{
			Id:   "p2o-ho-manager-solo",
			Name: "P2O (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Manager",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melaksanakan inovasi sistem pelayanan",
				"Melaksanakan dan mengevaluasi uji coba alat-alat penunjang pelayanan",
				"Merancang lay out pelayanan agar mampun mengingkatkan kinerja yang efektif",
				"Merancang program promosi di area pelayanan agar dapat meningkatkan penjualan",
			},
		},
		{
			Id:   "p2o-ho-supervisor-solo",
			Name: "P2O (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
		},
		{
			Id:   "catering-ho-manager-solo",
			Name: "Catering (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Manager",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
		},
		{
			Id:   "catering-ho-supervisor-solo",
			Name: "Catering (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
		},
		{
			Id:   "produksi-ho-operational-manager",
			Name: "Produksi (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Manager",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menerapkan sistem produksi yang efektif dan efisien",
				"Melakukan inovasi peralatan produksi agar lebih produktif",
				"Merancang lay out produksi agar mampun mengingkatkan kinerja yang efektif",
				"Melakukan analisa standar ritme produksi outlet untuk dapat memaksimalkan penjualan",
				"Merencanakan sistem pelaporan bagian produksi yang akurat",
			},
		},
		{
			Id:   "general-affair-area-bali",
			Name: "General Affair Area (Bali)",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Bali",
			},
			JobDesc: []string{
				"Memonitoring pembangunan/renovasi outlet",
				"Melakukan pengadaan dan perawatan perlengkapan outlet",
				"Mengelola keamanan outlet (dari pencurian, kebakaran, serangan hama)",
				"Mengelola dan merawat kendaraan outlet",
				"Mengelola database aset outlet",
			},
		},
		{
			Id:   "general-affair-area-surabaya",
			Name: "General Affair Area (Surabaya)",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Surabaya",
			},
			JobDesc: []string{
				"Memonitoring pembangunan/renovasi outlet",
				"Melakukan pengadaan dan perawatan perlengkapan outlet",
				"Mengelola keamanan outlet (dari pencurian, kebakaran, serangan hama)",
				"Mengelola dan merawat kendaraan outlet",
				"Mengelola database aset outlet",
			},
		},
		{
			Id:   "general-affair-area-pantura",
			Name: "General Affair Area (Pantura)",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Pantura",
			},
			JobDesc: []string{
				"Memonitoring pembangunan/renovasi outlet",
				"Melakukan pengadaan dan perawatan perlengkapan outlet",
				"Mengelola keamanan outlet (dari pencurian, kebakaran, serangan hama)",
				"Mengelola dan merawat kendaraan outlet",
				"Mengelola database aset outlet",
			},
		},
		{
			Id:   "general-affair-area-sumatra",
			Name: "General Affair Area (Sumatra)",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Sumatra",
			},
			JobDesc: []string{
				"Memonitoring pembangunan/renovasi outlet",
				"Melakukan pengadaan dan perawatan perlengkapan outlet",
				"Mengelola keamanan outlet (dari pencurian, kebakaran, serangan hama)",
				"Mengelola dan merawat kendaraan outlet",
				"Mengelola database aset outlet",
			},
		},
		{
			Id:   "finance-area-bali",
			Name: "Finance Area (Bali)",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Bali",
			},
			JobDesc: []string{
				"Memastikan kesesuaian dan akurasi seluruh laporan harian outlet di area masing-masing.",
				"Melakukan koordinasi dengan penanggung jawab laporan di setiap departemen atau outlet",
				"Melakukan validasinya laporan Cash Opnamedan laporan Stock Opname semua Outlet di Areanya",
				"Melakukan verifikasi secara fisik dokumen bukti transaksi / pembayaran",
				"Melakukan pendampingan terhadap Admin Outlet",
			},
		},
		{
			Id:   "finance-area-surabaya",
			Name: "Finance Area (Surabaya)",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Surabaya",
			},
			JobDesc: []string{
				"Memastikan kesesuaian dan akurasi seluruh laporan harian outlet di area masing-masing.",
				"Melakukan koordinasi dengan penanggung jawab laporan di setiap departemen atau outlet",
				"Melakukan validasinya laporan Cash Opnamedan laporan Stock Opname semua Outlet di Areanya",
				"Melakukan verifikasi secara fisik dokumen bukti transaksi / pembayaran",
				"Melakukan pendampingan terhadap Admin Outlet",
			},
		},
		{
			Id:   "finance-area-pantura",
			Name: "Finance Area (Pantura)",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "2",
				Penempatan: "Pantura",
			},
			JobDesc: []string{
				"Memastikan kesesuaian dan akurasi seluruh laporan harian outlet di area masing-masing.",
				"Melakukan koordinasi dengan penanggung jawab laporan di setiap departemen atau outlet",
				"Melakukan validasinya laporan Cash Opnamedan laporan Stock Opname semua Outlet di Areanya",
				"Melakukan verifikasi secara fisik dokumen bukti transaksi / pembayaran",
				"Melakukan pendampingan terhadap Admin Outlet",
			},
		},
		{
			Id:   "finance-area-jakarta",
			Name: "Finance Area (Jakarta)",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "3",
				Penempatan: "Jakarta",
			},
			JobDesc: []string{
				"Memastikan kesesuaian dan akurasi seluruh laporan harian outlet di area masing-masing.",
				"Melakukan koordinasi dengan penanggung jawab laporan di setiap departemen atau outlet",
				"Melakukan validasinya laporan Cash Opnamedan laporan Stock Opname semua Outlet di Areanya",
				"Melakukan verifikasi secara fisik dokumen bukti transaksi / pembayaran",
				"Melakukan pendampingan terhadap Admin Outlet",
			},
		},
		{
			Id:   "",
			Name: "",
			Infos: Info{
				Divisi:     "",
				Level:      "",
				Kebutuhan:  "",
				Penempatan: "",
			},
			JobDesc: []string{
				"",
				"",
				"",
				"",
				"",
			},
		},
	}
	for i := range spesifications {
		if spesifications[i].Infos.Divisi == "Human Capital" {
			spesifications[i].IsImage = true
			spesifications[i].Image = "../assets/images/human-capital.jpg"
		}
		if spesifications[i].Infos.Divisi == "Finance & Accounting" {
			spesifications[i].IsImage = true
			spesifications[i].Image = "../assets/images/finance-accounting.jpg"
		}
		if len(spesifications[i].JobDesc) > 0 {
			spesifications[i].IsJobDesc = true
		}
		if len(spesifications[i].QSpecific) > 0 {
			spesifications[i].IsSpesific = true
		}

		spesifications[i].QGeneric = []string{
			"Usia maksimal 35 tahun",
			"Pendidikan minimal S1 semua jurusan",
			"IPK minimal 3.0 (swasta) atau 2.75 (negeri)",
			"Penempatan: Solo",
			"Mampu bekerja dengan team maupun individu",
		}
	}
	c.HTML(http.StatusOK, "index.html", spesifications)
}
