@extends('layouts_dashboard.app')
@section('content')
  <!-- Content -->
  <div class="container-xxl flex-grow-1 container-p-y">
    <div class="card mb-4">
      <h5 class="card-header">Input Data {{ $title }}</h5>
      <form action="javascript:void(0)" id="form_data" name="form_data" method="POST" enctype="multipart/form-data"
        class="card-body">
        @csrf
        <div class="row g-3">
          <div class="row g-2">
            <div class="col-md-6">
              <label class="form-label" for="no_kk">Nomor Kartu Keluarga</label>
              <input type="text" id="no_kk" name="no_kk" class="form-control" placeholder="john.doe" />
              <span class="invalid-feedback" id="no_kk_error"></span>
            </div>
            <div class="col-md-6">
              <label class="form-label" for="nama_kepala_keluarga">Nama Kepala Keluarga</label>
              <input type="text" id="nama_kepala_keluarga" name="nama_kepala_keluarga" class="form-control"
                placeholder="john.doe" />
              <span class="invalid-feedback" id="nama_kepala_keluarga_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="id_provinsi">Provinsi</label>
              <select id="id_provinsi" name="id_provinsi" onchange="getKabupaten()" class="select2 form-select"
                data-allow-clear="true">
              </select>
              <span class="invalid-feedback" id="id_provinsi_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="id_kabupaten">Kabupaten</label>
              <select id="id_kabupaten" name="id_kabupaten" onchange="getKecamatan()" disabled class="select2 form-select"
                data-allow-clear="true">
              </select>
              <span class="invalid-feedback" id="id_kabupaten_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="id_kecamatan">Kecamatan</label>
              <select id="id_kecamatan" name="id_kecamatan" onchange="getDesa()" disabled class="select2 form-select"
                data-allow-clear="true">
              </select>
              <span class="invalid-feedback" id="id_kecamatan_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="id_desa">Desa</label>
              <select id="id_desa" name="id_desa" class="select2 form-select" disabled data-allow-clear="true">
              </select>
              <span class="invalid-feedback" id="id_desa_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="rt">RT</label>
              <input type="text" id="rt" name="rt" class="form-control" placeholder="john.doe" />
              <span class="invalid-feedback" id="rt_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="rw">RW</label>
              <input type="text" id="rw" name="rw" class="form-control" placeholder="john.doe" />
              <span class="invalid-feedback" id="rw_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="kode_pos">Kode Pos</label>
              <input type="text" id="kode_pos" name="kode_pos" class="form-control" placeholder="john.doe" />
              <span class="invalid-feedback" id="kode_pos_error"></span>
            </div>
            <div class="col-md-3">
              <label class="form-label" for="alamat">Alamat</label>
              <input type="text" id="alamat" name="alamat" class="form-control" placeholder="john.doe" />
              <span class="invalid-feedback" id="alamat_error"></span>
            </div>
          </div>
          <div class="pt-4">
            <button type="submit" id="saveBtn" value="create" class="btn btn-primary me-sm-3 me-1">Submit</button>
            <button type="reset" class="btn btn-label-secondary">Cancel</button>
          </div>
      </form>
    </div>
  </div>
  <!-- / Content -->
  @push('script')
    <script>
      $(document).ready(function() {
        fillProvinsiDropdown();
      });

      function fillProvinsiDropdown() {
        $.ajax({
          url: '/api/provinsi', // Ganti dengan URL API yang sesuai
          method: 'GET',
          success: function(data) {
            const provinsiDropdown = $('#id_provinsi');
            provinsiDropdown.empty();
            provinsiDropdown.append('<option value="">Pilih Provinsi</option>');

            const datas = data.data
            // Isi dropdown provinsi dengan data dari API
            datas.forEach(function(provinsi) {
              provinsiDropdown.append(`<option value="${provinsi.id}">${provinsi.nama_provinsi}</option>`);
            });
          },
          error: function(error) {
            console.error('Error fetching provinsi:', error.responseText);
          }
        });
      }

      function getKabupaten() {
        const selectedProvinsi = $('#id_provinsi').val();

        if (selectedProvinsi) {
          $.ajax({
            url: `/api/${selectedProvinsi}/kabupaten`, // Updated URL without double curly braces
            method: 'GET',
            success: function(data) {
              const kabupatenDropdown = $('#id_kabupaten');
              kabupatenDropdown.empty();
              kabupatenDropdown.append('<option value="">Pilih Kabupaten/Kota</option>');

              // Isi dropdown kabupaten dengan data dari API
              const datas = data.data;
              datas.forEach(function(kabupaten) {
                kabupatenDropdown.append(`<option value="${kabupaten.id}">${kabupaten.nama_kabkota}</option>`);
              });
              kabupatenDropdown.prop('disabled', false);
            },
            error: function(error) {
              console.error('Error fetching kabupaten:', error.responseText);
            }
          });
        } else {
          // Kosongkan dropdown kabupaten jika provinsi tidak dipilih
          $('#kabupaten').empty().append('<option value="">Pilih Kabupaten/Kota</option>');
        }
      }


      function getKecamatan() {
        const selectedKabupaten = $('#id_kabupaten').val();

        if (selectedKabupaten) {
          $.ajax({
            url: `/api/${selectedKabupaten}/kecamatan`, // Ganti dengan URL API yang sesuai
            method: 'GET',
            success: function(data) {
              const kecamatanDropdown = $('#id_kecamatan');
              kecamatanDropdown.empty();
              kecamatanDropdown.append('<option value="">Pilih Kecamatan</option>');

              // Isi dropdown kabupaten dengan data dari API
              const datas = data.data
              datas.forEach(function(kecamatan) {
                kecamatanDropdown.append(`<option value="${kecamatan.id}">${kecamatan.nama_kecamatan}</option>`);
              });
              kecamatanDropdown.prop('disabled', false);
            },
            error: function(error) {
              console.error('Error fetching kecamatan:', error.responseText);
            }
          });
        } else {
          // Kosongkan dropdown kabupaten jika provinsi tidak dipilih
          $('#id_kecamatan').empty().append('<option value="">Pilih Kecamatan</option>');
        }
      }

      function getDesa() {
        const selectedKecamatan = $('#id_kecamatan').val();

        if (selectedKecamatan) {
          $.ajax({
            url: `/api/${selectedKecamatan}/desa`, // Ganti dengan URL API yang sesuai
            method: 'GET',
            success: function(data) {
              const desaDropdown = $('#id_desa');
              const kodePosInput = $('#kode_pos');
              desaDropdown.empty();
              desaDropdown.append('<option value="">Pilih Desa</option>');

              // Isi dropdown kabupaten dengan data dari API
              const datas = data.data
              datas.forEach(function(desa) {
                desaDropdown.append(
                  `<option value="${desa.id}" data-kodepos="${desa.kodepos}">${desa.kelurahan}</option>`);
              });
              desaDropdown.prop('disabled', false);

              desaDropdown.change(function() {
                const selectedDesa = $(this).find(':selected');
                const kodePos = selectedDesa.data('kodepos');
                kodePosInput.val(kodePos);
              });
            },
            error: function(error) {
              console.error('Error fetching kabupaten:', error.responseText);
            }
          });
        } else {
          // Kosongkan dropdown kabupaten jika provinsi tidak dipilih
          $('#id_desa').empty().append('<option value="">Pilih Kabupaten/Kota</option>');
          $('#kode_pos').val('');
        }
      }


      $('#saveBtn').click(function(e) {
        e.preventDefault();
        $(this).html('Sending..');

        // Remove the error handling for the "jenis" and "Nama" fields
        $('#jenis').removeClass('is-invalid');
        $('#jenis-error').remove();

        var formData = new FormData($('#form_data')[0]);

        $.ajax({
          data: formData,
          url: "/api/kartu-keluarga",
          type: "POST",
          dataType: 'json',
          contentType: false,
          processData: false,
          success: function(data) {
            $('#form_data').trigger("reset");
            $('#saveBtn').html('Simpan');
            $('#ajaxModel').modal('hide');
            console.log(data)
            window.location.href = '{{ route('index.kartu_keluarga') }}';
            if (data.status) {
              Swal.fire('Sukses', data.message, 'success');
            } else {
              Swal.fire('Gagal', data.message, 'error');
            }
          },
          error: function(data) {
            console.log(data);
            $('#saveBtn').html('Save Changes');

            // Error handling for specific input fields
            if (data.responseJSON.errors) {
              var errors = data.responseJSON.errors;
              $.each(errors, function(key, value) {
                $("#" + key).addClass("is-invalid");
                $("#" + key + "_error").text(value[0]);
              });
            }
          }
        });
      });
    </script>
  @endpush
@endsection
