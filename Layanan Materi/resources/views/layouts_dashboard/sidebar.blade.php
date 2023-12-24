<aside id="layout-menu" class="layout-menu menu-vertical menu bg-menu-theme">
    <div class="app-brand demo">
      <a href="index.html" class="app-brand-link">
        <span class="app-brand-logo demo">
          <svg width="32" height="22" viewBox="0 0 32 22" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M0.00172773 0V6.85398C0.00172773 6.85398 -0.133178 9.01207 1.98092 10.8388L13.6912 21.9964L19.7809 21.9181L18.8042 9.88248L16.4951 7.17289L9.23799 0H0.00172773Z"
              fill="#7367F0" />
            <path
              opacity="0.06"
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M7.69824 16.4364L12.5199 3.23696L16.5541 7.25596L7.69824 16.4364Z"
              fill="#161616" />
            <path
              opacity="0.06"
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M8.07751 15.9175L13.9419 4.63989L16.5849 7.28475L8.07751 15.9175Z"
              fill="#161616" />
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M7.77295 16.3566L23.6563 0H32V6.88383C32 6.88383 31.8262 9.17836 30.6591 10.4057L19.7824 22H13.6938L7.77295 16.3566Z"
              fill="#7367F0" />
          </svg>
        </span>
        <span class="app-brand-text demo menu-text fw-bold">SIMAMI</span>
      </a>

      <a href="javascript:void(0);" class="layout-menu-toggle menu-link text-large ms-auto">
        <i class="ti menu-toggle-icon d-none d-xl-block ti-sm align-middle"></i>
        <i class="ti ti-x d-block d-xl-none ti-sm align-middle"></i>
      </a>
    </div>

    <div class="menu-inner-shadow"></div>

    <ul class="menu-inner py-1">
      <!-- Dashboards -->
      @if(in_array('dashboard', $userAbilities))
      <li class="menu-item {{ (request()->is('dashboard*')) ? 'active' : '' }}">
        <a href="{{ route('dashboard') }}" class="menu-link">
          <i class="menu-icon tf-icons ti ti-smart-home"></i>
          <div data-i18n="Dashboards">Dashboards</div>
        </a>
      </li>
      @endif
      @if(in_array('master-data', $userAbilities))
      <li class="menu-header small text-uppercase">
        <span class="menu-header-text">Master</span>
      </li>
      <li class="menu-item {{ (request()->is('m-wilayah*')) ? 'active open' : '' }}">
        <a href="javascript:void(0);" class="menu-link menu-toggle">
          <i class="menu-icon tf-icons ti ti ti-world"></i>
          <div data-i18n="Wilayah">Wilayah</div>
        </a>
        <ul class="menu-sub">
          <li class="menu-item {{ (request()->is('m-wilayah/provinsi*')) ? 'active' : '' }}">
            <a href="/m-wilayah/provinsi" class="menu-link">
              <div data-i18n="Provinsi">Provinsi</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('m-wilayah/kabupaten*')) ? 'active' : '' }}">
            <a href="/m-wilayah/kabupaten" class="menu-link">
              <div data-i18n="Kabupaten">Kabupaten</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('wilayah/kecamatan*')) ? 'active' : '' }}">
            <a href="/wilayah/kecamatan" class="menu-link">
              <div data-i18n="Kecamatan">Kecamatan</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('wilayah/desa*')) ? 'active' : '' }}">
            <a href="/wilayah/desa" class="menu-link">
              <div data-i18n="Desa">Desa</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('wilayah/banjar*')) ? 'active' : '' }}">
            <a href="/wilayah/banjar" class="menu-link">
              <div data-i18n="Banjar">Banjar</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('wilayah/tempek*')) ? 'active' : '' }}">
            <a href="/wilayah/tempek" class="menu-link">
              <div data-i18n="Tempek">Tempek</div>
            </a>
          </li>
        </ul>
      </li>
      <li class="menu-item {{ (request()->is('m-user*')) ? 'active open' : '' }}">
        <a href="javascript:void(0);" class="menu-link menu-toggle">
          <i class="menu-icon tf-icons ti ti-user"></i>
          <div data-i18n="Manajemen User">Manajemen User</div>
        </a>
        <ul class="menu-sub">
          <li class="menu-item {{ (request()->is('m-user/role*')) ? 'active' : '' }}">
            <a href="/m-user/role" class="menu-link">
              <div data-i18n="Role">Role</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('m-user/permission*')) ? 'active' : '' }}">
            <a href="/m-user/permission" class="menu-link">
              <div data-i18n="Permission">Permission</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('m-user/user-account*')) ? 'active' : '' }}">
            <a href="/m-user/user-account" class="menu-link">
              <div data-i18n="Manajemen User">Manajemen User</div>
            </a>
          </li>
        </ul>
      </li>
      <li class="menu-item {{ (request()->is('m-iuran*')) ? 'active open' : '' }}">
        <a href="javascript:void(0);" class="menu-link menu-toggle">
          <i class="menu-icon tf-icons ti ti-file-dollar"></i>
          <div data-i18n="Iuran">Iuran</div>
        </a>
        <ul class="menu-sub">
          <li class="menu-item {{ (request()->is('m-iuran/tipe-iuran*')) ? 'active' : '' }}">
            <a href="/m-iuran/tipe-iuran" class="menu-link">
              <div data-i18n="Tipe Iuran">Tipe Iuran</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('m-iuran/syarat-iuran*')) ? 'active' : '' }}">
            <a href="/m-iuran/syarat-iuran" class="menu-link">
              <div data-i18n="Syarat Iuran">Syarat Iuran</div>
            </a>
          </li>
        </ul>
      </li>
      <li class="menu-item {{ (request()->is('m-dokumen*')) ? 'active open' : '' }}">
        <a href="javascript:void(0);" class="menu-link menu-toggle">
          <i class="menu-icon tf-icons ti ti-clipboard-text"></i>
          <div data-i18n="Dokumen">Dokumen</div>
        </a>
        <ul class="menu-sub">
          <li class="menu-item {{ (request()->is('m-dokumen/jenis-dokumen*')) ? 'active' : '' }}">
            <a href="/m-dokumen/jenis-dokumen" class="menu-link">
              <div data-i18n="Tipe Dokumen">Tipe Dokumen</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('m-dokumen/syarat-dokumen*')) ? 'active' : '' }}">
            <a href="/m-dokumen/syarat-dokumen" class="menu-link">
              <div data-i18n="Syarat Dokumen">Syarat Dokumen</div>
            </a>
          </li>
        </ul>
      </li>
      <li class="menu-item {{ (request()->is('m-master-pencatatan*')) ? 'active open' : '' }}">
        <a href="javascript:void(0);" class="menu-link menu-toggle">
          <i class="menu-icon tf-icons ti ti-file-certificate"></i>
          <div data-i18n="Pencatatan">Pencatatan</div>
        </a>
        <ul class="menu-sub">
          <li class="menu-item {{ (request()->is('m-master-pencatatan/jenis-pekerjaan*')) ? 'active' : '' }}">
            <a href="/m-master-pencatatan/jenis-pekerjaan" class="menu-link">
              <div data-i18n="Jenis Pekerjaan">Jenis Pekerjaan</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('m-master-pencatatan/jenis-pendidikan*')) ? 'active' : '' }}">
            <a href="/m-master-pencatatan/jenis-pendidikan" class="menu-link">
              <div data-i18n="Jenis Pendidikan">Jenis Pendidikan</div>
            </a>
          </li>
        </ul>
      </li>
      <li class="menu-item {{ (request()->is('m-kategori-pengumuman*')) ? 'active' : '' }}">
        <a href="/m-kategori-pengumuman" class="menu-link">
          <i class="menu-icon tf-icons ti ti-notes"></i>
          <div data-i18n="Kategori Pengumuman">Kategori Pengumuman</div>
        </a>
      </li>
      @endif
      @if(in_array('pencatatan-penduduk', $userAbilities))
      <!-- Apps & Pages -->
      <li class="menu-header small text-uppercase">
        <span class="menu-header-text">Penduduk</span>
      </li>
      <li class="menu-item {{ (request()->is('penduduk/kartu-keluarga*')) ? 'active' : '' }}">
        <a href="/penduduk/kartu-keluarga" class="menu-link">
          <i class="menu-icon tf-icons ti ti-users"></i>
          <div data-i18n="Kartu Keluarga">Kartu Keluarga</div>
        </a>
      </li>
      <li class="menu-item {{ (request()->is('penduduk/anggota-keluarga*')) ? 'active' : '' }}">
        <a href="/penduduk/anggota-keluarga" class="menu-link">
          <i class="menu-icon tf-icons ti ti-id"></i>
          <div data-i18n="Penduduk">Penduduk</div>
        </a>
      </li>
      <li class="menu-item {{ (request()->is('penduduk/mutasi*')) ? 'active open' : '' }}">
        <a href="javascript:void(0);" class="menu-link menu-toggle">
          <i class="menu-icon tf-icons ti ti-file-certificate"></i>
          <div data-i18n="Mutasi">Mutasi</div>
        </a>
        <ul class="menu-sub">
          <li class="menu-item {{ (request()->is('penduduk/mutasi/kelahiran*')) ? 'active' : '' }}">
            <a href="/penduduk/mutasi/kelahiran" class="menu-link">
              <div data-i18n="Kelahiran">Kelahiran</div>
            </a>
          </li>
          <li class="menu-item {{ (request()->is('penduduk/mutasi/kematian*')) ? 'active' : '' }}">
            <a href="/penduduk/mutasi/kematian" class="menu-link">
              <div data-i18n="Kematian">Kematian</div>
            </a>
          </li>
        </ul>
      </li>
      @endif
      @if(in_array('iuran', $userAbilities))
      <li class="menu-header small text-uppercase">
        <span class="menu-header-text">Iuran</span>
      </li>
      <li class="menu-item">
        <a href="app-email.html" class="menu-link">
          <i class="menu-icon tf-icons ti ti-cash-banknote"></i>
          <div data-i18n="Buat Iuran">Buat Iuran</div>
        </a>
      </li>
      <li class="menu-item">
        <a href="app-email.html" class="menu-link">
          <i class="menu-icon tf-icons ti ti-report-money"></i>
          <div data-i18n="History Pembayaran">History Pembayaran</div>
        </a>
      </li>
      @endif
      @if(in_array('dokumen', $userAbilities))
      <li class="menu-header small text-uppercase">
        <span class="menu-header-text">Dokumen</span>
      </li>
      <li class="menu-item">
        <a href="app-email.html" class="menu-link">
          <i class="menu-icon tf-icons ti ti-file-description"></i>
          <div data-i18n="Pengajuan Dokumen">Pengajuan Dokumen</div>
        </a>
      </li>
      @endif
      @if(in_array('pengumuman', $userAbilities))
      <li class="menu-header small text-uppercase">
        <span class="menu-header-text">Pengaduan</span>
      </li>
      <li class="menu-item">
        <a href="app-email.html" class="menu-link">
          <i class="menu-icon tf-icons ti ti-lifebuoy"></i>
          <div data-i18n="List Pengaduan">List Pengaduan</div>
        </a>
      </li>
      @endif
    </ul>
  </aside>