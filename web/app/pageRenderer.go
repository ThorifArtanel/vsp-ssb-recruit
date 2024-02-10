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
			Id:   "internalaudit-seniormanager-bugged",
			Name: "Internal Audit",
			Infos: Info{
				Divisi:     "Internal Audit",
				Level:      "Senior Manager",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "cultureorganizationaldevelopment-manager-bugged",
			Name: "Culture Organizational Development",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Manager",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:       "recruitmentandselection-staff-solo",
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
			Id:   "ssbcorpuniversity-staff-solo",
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
			Id:   "sotfwareengineer-staff-solo",
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
			Id:   "p2oho-manager-solo",
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
			Id:   "p2oho-supervisor-solo",
			Name: "P2O (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
		},
		{
			Id:   "cateringho-manager-solo",
			Name: "Catering (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Manager",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
		},
		{
			Id:   "cateringho-supervisor-solo",
			Name: "Catering (HO)",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
		},
		{
			Id:   "produksiho-manager-solo",
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
			Id:   "manageroutlet-supervisor-bugged",
			Name: "Manager Outlet",
			Infos: Info{
				Divisi:     "Operational",
				Level:      "Supervisor",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "generalaffairarea-staff-bali",
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
			Id:   "generalaffairarea-staff-surabaya",
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
			Id:   "generalaffairarea-staff-pantura",
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
			Id:   "generalaffairarea-staff-sumatra",
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
			Id:   "financearea-staff-bali",
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
			Id:   "financearea-staff-surabaya",
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
			Id:   "financearea-staff-pantura",
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
			Id:   "financearea-staff-jakarta",
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
			Id:   "accounting-supervisor-solo",
			Name: "Accounting",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menyediakan data dan laporan keuangan yang akurat",
				"Menyusun laporan keuangan sesuai dengan PSAK untuk mengetahui kondisi keuangan perusahaan",
				"Membuat analisa laporan keuangan",
				"Menyusun laporan evaluasi analisa keuangan untuk bagian terkait",
			},
		},
		{
			Id:   "accountpayablecollector-staff-solo",
			Name: "Account Payable Collector",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "2",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Memeriksa alur dan kelengkapan berkas pengajuan anggaran",
				"Menyerahkan berkas pengajuan anggaran kepada AP Verificator",
				"Menerima dan mengirimkan bukti transfer pengajuan anggaran kepada pihak terkait",
				"Mengumpulkan laporan penggunaan/realisasi anggaran",
				"Mengontrol penggunaan anggaran Petty Cash operasional di seluruh Outlet",
			},
		},
		{
			Id:   "bookkeeper-staff-solo",
			Name: "Book Keeper",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melaksanakan dan menyusun laporan Cashopname Bank",
				"Menyusun Laporan Realisasi RAB outlet baru",
				"Menyusun Laporan Harta Outlet",
				"Menyusun Laporan Laba Rugi Outlet Kepada Mitra",
				"Melakukan pengarsipan dokumen berjalan sesuai degan ketentuan",
			},
		},
		{
			Id:   "generalaccountingbank-staff-solo",
			Name: "General Accounting Bank",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menyusun Laporan Cashopname Bank",
				"Melakukan penjurnalan semua data transaksi sesuai ketentuan perusahaan",
				"Menyusun Laporan Petty Cash",
				"Mengumpulkan data dan verifikasi transaksi petty cash outlet",
				"Mengecek dan memastikan dokumen transaksi sesuai dengan SOP",
				"Mengarsipkan dokument transaksi berdasarkan masing-masing outlet",
			},
		},
		{
			Id:   "staffarsip-staff-solo",
			Name: "Staff Arsip",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melakukan pengelolaan arsip dinamis maupun statis",
				"Mengumpulkan seluruh bukti transaksi/dokumen penting lain secara periodik",
				"Melakukan pengolahan dan penyajian arsip menjadi informasi yang dibutuhkan",
				"Menjaga ketersediaan arsip yang autentik dan terpercaya sebagai alat bukti yang sah",
			},
		},
		{
			Id:   "staffcostaccounting-staff-solo",
			Name: "Staff Cost Accounting",
			Infos: Info{
				Divisi:     "Finance & Accounting",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melakukan perencanaan, mengatur dan mengontrol analisis keuangan",
				"Melakukan analisis laporan efisiensi produksi",
				"Membuatkan laporan harga pokok produksi",
				"Melakukan analisa harga bahan baku",
			},
		},
		{
			Id:   "recruitment&assessment-staff-solo",
			Name: "Recruitment & Assessment",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Membuat konten flyer (job vacancy) & job posting",
				"Melakukan screening dan seleksi calon karyawan",
				"Melaksanakan interview calon karyawan outlet",
				"Mengolah data dan hasil seleksi calon karyawan",
				"Melaksanakan proses assessment untuk promosi dan mutasi",
			},
		},
		{
			Id:   "learning&development-supervisor-solo",
			Name: "Learning & Development",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menganalisa kebutuhan pada pelatihan untuk meningkatkan keterampilan karyawan",
				"Merencanakan pelaksanaan training karyawan baru maupun existing",
				"Melakukan evaluasi hasil pelaksanaan training karyawan",
				"Merancang jenis dan metode pelatihan atau training sesuai kebutuhan perusahaan",
			},
		},
		{
			Id:   "learning&development-staff-solo",
			Name: "Learning & Development",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menyusun surat menyurat",
				"Melaksanakan pelatihan internal perusahaan",
				"Menjalankan training booth camp",
				"Menyusun laporan pelaksanaan training",
				"Menjalankan fungsi event organizer dalam acara internal perusahaan",
			},
		},
		{
			Id:   "hcarea-supervisor-bali",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Bali",
			},
			JobDesc: []string{
				"Menyusun dan mengevaluasi man power planning (MPP) outlet",
				"Melakukan seleksi calon karyawan baru",
				"Melakukan coaching, counseling dan mentoring karyawan",
				"Melaksanakan pelatihan internal perusahaan",
				"Memastikan penegakan peraturan perusahaan diseluruh outlet",
			},
		},
		{
			Id:   "hcarea-supervisor-surabaya",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Surabaya",
			},
			JobDesc: []string{
				"Menyusun dan mengevaluasi man power planning (MPP) outlet",
				"Melakukan seleksi calon karyawan baru",
				"Melakukan coaching, counseling dan mentoring karyawan",
				"Melaksanakan pelatihan internal perusahaan",
				"Memastikan penegakan peraturan perusahaan diseluruh outlet",
			},
		},
		{
			Id:   "hcarea-supervisor-pantura",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Pantura",
			},
			JobDesc: []string{
				"Menyusun dan mengevaluasi man power planning (MPP) outlet",
				"Melakukan seleksi calon karyawan baru",
				"Melakukan coaching, counseling dan mentoring karyawan",
				"Melaksanakan pelatihan internal perusahaan",
				"Memastikan penegakan peraturan perusahaan diseluruh outlet",
			},
		},
		{
			Id:   "hcarea-supervisor-sumatra",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Sumatra",
			},
			JobDesc: []string{
				"Menyusun dan mengevaluasi man power planning (MPP) outlet",
				"Melakukan seleksi calon karyawan baru",
				"Melakukan coaching, counseling dan mentoring karyawan",
				"Melaksanakan pelatihan internal perusahaan",
				"Memastikan penegakan peraturan perusahaan diseluruh outlet",
			},
		},
		{
			Id:   "hcarea-staff-solo",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menyusun surat menyurat",
				"Melakukan seleksi calon karyawan baru",
				"Memastikan penyerahan benefit seluruh karyawan",
				"Melaksanakan pelatihan internal perusahaan",
			},
		},
		{
			Id:   "hcarea-staff-pantura",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Pantura",
			},
			JobDesc: []string{
				"Menyusun surat menyurat",
				"Melakukan seleksi calon karyawan baru",
				"Memastikan penyerahan benefit seluruh karyawan",
				"Melaksanakan pelatihan internal perusahaan",
			},
		},
		{
			Id:   "hcarea-staff-bandung",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Bandung",
			},
			JobDesc: []string{
				"Menyusun surat menyurat",
				"Melakukan seleksi calon karyawan baru",
				"Memastikan penyerahan benefit seluruh karyawan",
				"Melaksanakan pelatihan internal perusahaan",
			},
		},
		{
			Id:   "hcarea-staff-jakarta",
			Name: "HC Area",
			Infos: Info{
				Divisi:     "Human Capital",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Jakarta",
			},
			JobDesc: []string{
				"Menyusun surat menyurat",
				"Melakukan seleksi calon karyawan baru",
				"Memastikan penyerahan benefit seluruh karyawan",
				"Melaksanakan pelatihan internal perusahaan",
			},
		},
		{
			Id:   "marketing-seniormanager-bugged",
			Name: "Marketing",
			Infos: Info{
				Divisi:     "Marketing",
				Level:      "Senior Manager",
				Kebutuhan:  "0",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "kontenkreator-supervisor-solo",
			Name: "Konten Kreator",
			Infos: Info{
				Divisi:     "Marketing",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menulis, meninjau, mengedit, dan membuat konten untuk platform yang digunakan perusahaan",
				"Melakukan riset dan interview untuk mempelajari tren terkini serta dalam pengembangan konten",
				"Mempersiapkan materi promosi",
				"Menggunakan media sosial untuk consumer engagement, merespons komplain, dan mempromosikan produk/layanan perusahaan",
				"Mengelola media sosial dan company website",
				"Meningkatkan traffic melalui konten yang dibuat",
			},
		},
		{
			Id:   "customerrelationmarketing-supervisor-solo",
			Name: "Customer Relation Marketing",
			Infos: Info{
				Divisi:     "Marketing",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Melacak prospek atau calon pembeli",
				"Melakukan follow up customer pasif",
				"Melakukan follow up untuk perbaikan pada produk yang dikomplain",
				"Mengelola hubungan baik dengan pelanggan",
			},
		},
		{
			Id:   "marketingbranding-staff-solo",
			Name: "Marketing Branding",
			Infos: Info{
				Divisi:     "Marketing",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Mengikuti dan mengorganisir keikutsertaan dalam event eksternal",
				"Melakukan pembaharuan materi branding outlet",
				"Melakukan optimalisasi sosial media",
				"Melakukan kerjasama dengan pihak eksternal untuk melakukan promosi",
			},
		},
		{
			Id:   "marketingsales-supervisor-solo",
			Name: "Marketing Sales",
			Infos: Info{
				Divisi:     "Marketing",
				Level:      "Supervisor",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Menyusun strategi penjualan produk",
				"Menganalisa pasar dan tren",
				"Mencari target konsumen",
				"Melakukan penawaran produk kepada konsumen/instansi",
				"Merekap dan menganalisa laporan hasil penjualan",
			},
		},
		{
			Id:   "surveyor-staff-solo",
			Name: "Surveyor",
			Infos: Info{
				Divisi:     "Business Development",
				Level:      "Staff",
				Kebutuhan:  "2",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Mencari area atau lokasi prospek secara online maupun offline",
				"Melakukan survei dan pemetaan lahan untuk keperluan opening outlet",
				"Menganalisa potensi area sekitar lokasi yang disurvei",
				"Membuat laporan dari hasil survei dan mempresentasikan kepada atasan",
			},
		},
		{
			Id:   "kemitraan-staff-solo",
			Name: "Kemitraan",
			Infos: Info{
				Divisi:     "Business Development",
				Level:      "Staff",
				Kebutuhan:  "1",
				Penempatan: "Solo",
			},
			JobDesc: []string{
				"Memberikan informasi perusahaan kepada investor secara persuasif",
				"Menjadwalkan sosialisasi sistem kerjasama dengan invenstor",
				"Melakukan sosialisasi sistem kerjasama kepada calon investor",
				"Menjaga hubungan baik perusahaan dengan investor",
			},
		},
		{
			Id:   "supplychainmanagement-seniormanager-bugged",
			Name: "Supply Chain Management",
			Infos: Info{
				Divisi:     "SCM",
				Level:      "Senior Manager",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Memberikan informasi perusahaan kepada investor secara persuasif",
				"Menjadwalkan sosialisasi sistem kerjasama dengan invenstor",
				"Melakukan sosialisasi sistem kerjasama kepada calon investor",
				"Menjaga hubungan baik perusahaan dengan investor",
			},
		},
		{
			Id:   "logisticarea-supervisor-bugged",
			Name: "Logistic Area",
			Infos: Info{
				Divisi:     "SCM",
				Level:      "Supervisor",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "purchasing-staff-bugged",
			Name: "Purchasing",
			Infos: Info{
				Divisi:     "Central Kitchen",
				Level:      "Staff",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "supplyplan-staff-bugged",
			Name: "Supply Plan",
			Infos: Info{
				Divisi:     "SCM",
				Level:      "Staff",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "warehouse-supervisor-bugged",
			Name: "Warehouse",
			Infos: Info{
				Divisi:     "SCM",
				Level:      "Supervisor",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
			},
		},
		{
			Id:   "maintenanceataumekanik-supervisor-bugged",
			Name: "Maintenance / Mekanik",
			Infos: Info{
				Divisi:     "Central Kitchen",
				Level:      "Supervisor",
				Kebutuhan:  "Bugged",
				Penempatan: "Bugged",
			},
			JobDesc: []string{
				"Bugged",
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
