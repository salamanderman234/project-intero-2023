@extends('layouts_dashboard.app')
@section('content')
  <!-- Content -->
  <div class="container-xxl flex-grow-1 container-p-y">
    <!-- Users List Table -->
    <div class="card">
      <div class="card-header d-flex align-items-center justify-content-between">
        <h5 class="card-title">List {{ $title }}</h5>
        <a href="{{ route('create.kartu_keluarga') }}" id="addNewData" class="btn btn-primary">
          <i class="ti ti-plus"></i>&nbsp; Tambah Data
        </a>
      </div>
      <div class="card-datatable table-responsive">
        <table class="datatables-tempek table border-top">
          <thead>
            <tr>
              <th>No</th>
              <th>No. KK</th>
              <th>Nama Kepala Keluarga</th>
              <th>Anggota Keluarga</th>
              <th>Actions</th>
            </tr>
          </thead>
        </table>
      </div>
      <!-- Modal -->
      <div class="modal fade" id="basicModal" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="exampleModalLabel1">Tambah Data Tempek</h5>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <div class="row">
                <div class="col mb-3">
                  <label for="nameBasic" class="form-label">Nama Tempek<span style="color: red">*</span></label>
                  <input type="text" id="nameBasic" class="form-control" placeholder="Masukkan nama Tempek" />
                </div>
              </div>
              <div class="row g-2">
                <div class="col mb-3">
                  <label for="emailBasic" class="form-label">Deskripsi</label>
                  <textarea id="autosize-demo" rows="3" class="form-control" placeholder="Masukkan deskripsi Tempek"></textarea>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-label-secondary" data-bs-dismiss="modal">
                Close
              </button>
              <button type="button" class="btn btn-primary">Save changes</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- / Content -->
  @push('script')
    <script>
      'use strict';

      var dt_user;
      // Datatable (jquery)
      $(document).ready(function() {
        let borderColor, bodyBg, headingColor;

        if (isDarkStyle) {
          borderColor = config.colors_dark.borderColor;
          bodyBg = config.colors_dark.bodyBg;
          headingColor = config.colors_dark.headingColor;
        } else {
          borderColor = config.colors.borderColor;
          bodyBg = config.colors.bodyBg;
          headingColor = config.colors.headingColor;
        }

        // Variable declaration for table
        var dt_user_table = $('.datatables-tempek'),
          select2 = $('.select2'),
          userView = 'app-user-view-account.html',
          statusObj = {
            1: {
              title: 'Pending',
              class: 'bg-label-warning'
            },
            2: {
              title: 'Active',
              class: 'bg-label-success'
            },
            3: {
              title: 'Inactive',
              class: 'bg-label-secondary'
            }
          };

        if (select2.length) {
          var $this = select2;
          $this.wrap('<div class="position-relative"></div>').select2({
            placeholder: 'Select Country',
            dropdownParent: $this.parent()
          });
        }

        // Users datatable
        if (dt_user_table.length) {
          dt_user = dt_user_table.DataTable({
            ajax: {
              url: '{{ url('api/kartu-keluarga') }}',
              type: "GET",
              dataSrc: 'data'
            },
            columns: [
              // columns according to JSON
              {
                data: ''
              },
              {
                data: 'no_kk'
              },
              {
                data: 'nama_kepala_keluarga'
              },
              {
                data: 'kecamatan'
              },
              {
                data: 'action'
              }
            ],
            columnDefs: [{
                // For Responsive
                className: 'control',
                searchable: false,
                orderable: false,
                responsivePriority: 2,
                targets: 0,
                render: function(data, type, full, meta) {
                  return meta.row + 1; // Menampilkan nomor urut sesuai urutan kolom
                }
              },
              {
                // Plans
                targets: 1,
                render: function(data, type, full, meta) {
                  var $no_kk = full['no_kk'];

                  return '<span class="fw-semibold">' + $no_kk + '</span>';
                }
              },
              {
                // Plans
                targets: 2,
                render: function(data, type, full, meta) {
                  var $nama_kepala_keluarga = full['nama_kepala_keluarga'];

                  return '<span class="fw-semibold">' + $nama_kepala_keluarga + '</span>';
                }
              },
              {
                // Teams
                targets: 3,
                orderable: false,
                searchable: false,
                render: function(data, type, full, meta) {
                  // For Avatar badge
                  var anggotaKeluargaHtml = '';
                  if (full['anggota_keluarga']) {
                    anggotaKeluargaHtml += '<div class="d-flex align-items-center avatar-group">';

                    full['anggota_keluarga'].forEach(function(penduduks) {
                      var stateNum = Math.floor(Math.random() * 6);
                      var states = ['success', 'danger', 'warning', 'info', 'primary', 'secondary'];
                      var $state = states[stateNum],
                        $name = penduduks['nama'],
                        $initials = $name.match(/\b\w/g) || [];
                      $initials = (($initials.shift() || '') + ($initials.pop() || '')).toUpperCase();

                      anggotaKeluargaHtml +=
                        '<div data-bs-toggle="tooltip" data-popup="tooltip-custom" data-bs-placement="top" title="' +
                        $name +
                        '" class="avatar"><span class="avatar-initial rounded-circle pull-up bg-label-' +
                        $state + '">' + $initials + '</span></div>';
                    });

                    anggotaKeluargaHtml += '</div>';
                  } else {
                    anggotaKeluargaHtml = '-';
                  }
                  return anggotaKeluargaHtml;
                }
              },
              {
                // Actions
                targets: -1,
                title: 'Actions',
                searchable: false,
                orderable: false,
                render: function(data, type, full, meta) {
                  var $id = full['id'];
                  return (
                    '<div class="d-flex align-items-center">' +
                    '<a href="/penduduk/' + $id +
                    '/anggota-keluarga" class="text-body" data-bs-toggle="tooltip" data-bs-placement="top" title="Lihat Anggota Keluarga"><i class="ti ti-user ti-sm me-2"></i></a>' +
                    '<a href="/penduduk/kartu-keluarga/' + $id +
                    '" class="text-body" data-bs-toggle="tooltip" data-bs-placement="top" title="Edit Kartu Keluarga"><i class="ti ti-edit ti-sm mx-2"></i></a>' +
                    '<a href="javascript:void(0)" onclick="deleteData(' + $id +
                    ')" class="text-body" data-id=' + $id +
                    ' data-bs-toggle="tooltip" data-bs-placement="top" title="Hapus Kartu Keluarga" delete-record"><i class="ti ti-trash ti-sm mx-2"></i></a>' +
                    '</div>'
                  );
                }
              }
            ],
            order: [],
            dom: '<"row me-2"' +
              '<"col-md-2"<"me-3"l>>' +
              '<"col-md-10"<"dt-action-buttons text-xl-end text-lg-start text-md-end text-start d-flex align-items-center justify-content-end flex-md-row flex-column mb-3 mb-md-0"fB>>' +
              '>t' +
              '<"row mx-2"' +
              '<"col-sm-12 col-md-6"i>' +
              '<"col-sm-12 col-md-6"p>' +
              '>',
            language: {
              sLengthMenu: '_MENU_',
              search: '',
              searchPlaceholder: 'Search..'
            },
            // Buttons with Dropdown
            buttons: [{
              extend: 'collection',
              className: 'btn btn-label-primary dropdown-toggle mx-3',
              text: '<i class="ti ti-screen-share me-1 ti-xs"></i>Export',
              buttons: [{
                  extend: 'print',
                  text: '<i class="ti ti-printer me-2" ></i>Print',
                  className: 'dropdown-item',
                  exportOptions: {
                    columns: [1, 2, 3, 4, 5],
                    // prevent avatar to be print
                    format: {
                      body: function(inner, coldex, rowdex) {
                        if (inner.length <= 0) return inner;
                        var el = $.parseHTML(inner);
                        var result = '';
                        $.each(el, function(index, item) {
                          if (item.classList !== undefined && item.classList.contains('user-name')) {
                            result = result + item.lastChild.firstChild.textContent;
                          } else if (item.innerText === undefined) {
                            result = result + item.textContent;
                          } else result = result + item.innerText;
                        });
                        return result;
                      }
                    }
                  },
                  customize: function(win) {
                    //customize print view for dark
                    $(win.document.body)
                      .css('color', headingColor)
                      .css('border-color', borderColor)
                      .css('background-color', bodyBg);
                    $(win.document.body)
                      .find('table')
                      .addClass('compact')
                      .css('color', 'inherit')
                      .css('border-color', 'inherit')
                      .css('background-color', 'inherit');
                  }
                },
                {
                  extend: 'excel',
                  text: '<i class="ti ti-file-spreadsheet me-2"></i>Excel',
                  className: 'dropdown-item',
                  exportOptions: {
                    columns: [1, 2, 3, 4, 5],
                    // prevent avatar to be display
                    format: {
                      body: function(inner, coldex, rowdex) {
                        if (inner.length <= 0) return inner;
                        var el = $.parseHTML(inner);
                        var result = '';
                        $.each(el, function(index, item) {
                          if (item.classList !== undefined && item.classList.contains('user-name')) {
                            result = result + item.lastChild.firstChild.textContent;
                          } else if (item.innerText === undefined) {
                            result = result + item.textContent;
                          } else result = result + item.innerText;
                        });
                        return result;
                      }
                    }
                  }
                },
              ]
            }, ],
          });
        }

        // $('.dataTables_filter').html('<div class="input-group flex-nowrap"><span class="input-group-text" id="addon-wrapping"><i class="ti ti-search"></i></span><input type="search" class="form-control form-control-sm" placeholder="Type in to Search" aria-label="Type in to Search" aria-describedby="addon-wrapping"></div>');

        // Filter form control to default size
        // ? setTimeout used for multilingual table initialization
        setTimeout(() => {
          $('.dataTables_filter .form-control').removeClass('form-control-sm');
          $('.dataTables_length .form-select').removeClass('form-select-sm');
        }, 300);
      });

      // Function to delete data
      function deleteData(id) {
        Swal.fire({
          icon: 'warning',
          text: 'Hapus Data Jenis Museum?',
          showCancelButton: true,
          confirmButtonText: 'Hapus',
          cancelButtonText: 'Batal',
        }).then((result) => {
          if (result.isConfirmed) {
            $.ajax({
              url: "/api/kartu-keluarga/" + id,
              type: "DELETE",
              dataType: "JSON",
              success: function(data) {
                if (data.status) {
                  Swal.fire('Sukses', data.message, 'success');
                  // Reload the DataTable after successful deletion
                  dt_user.ajax.reload();
                } else {
                  Swal.fire('Gagal', data.message, 'error');
                }
              },
              error: function(error) {
                Swal.fire('Gagal', 'terjadi kesalahan sistem', 'error');
                console.log(error.XMLHttpRequest);
              }
            });
          }
        });
      }
    </script>
  @endpush
@endsection
