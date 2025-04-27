<script lang="ts">
  import { goto } from '$app/navigation';

  let nama: string = '';
  let tanggalLahir: string = '';
  let jenisKelamin: string = '';
  let pekerjaan: string = '';
  let foto: FileList | undefined;

  async function handleSubmit(event: SubmitEvent): Promise<void> {
    event.preventDefault();

    const formData = new FormData();
    formData.append('nama', nama);
    formData.append('tanggal_lahir', tanggalLahir);
    formData.append('jenis_kelamin', jenisKelamin);
    formData.append('pekerjaan', pekerjaan);

    if (foto && foto.length > 0) {
      formData.append('foto', foto[0]);
    }

    try {
      const response = await fetch('http://localhost:8080/users', {
        method: 'POST',
        body: formData
      });

      if (response.ok) {
        alert('User berhasil disimpan!');
        goto('/'); // kembali ke halaman utama atau tabel
      } else {
        const err = await response.text();
        alert('Gagal menyimpan: ' + err);
      }
    } catch (error) {
      alert('Terjadi kesalahan saat menyimpan data.');
      console.error(error);
    }
  }
</script>

<!-- UI-nya tetap sama seperti sebelumnya -->
<div class="container mt-5">
  <div class="card shadow">
    <div class="card-header bg-primary text-white">
      <h5 class="mb-0"><i class="bi bi-person-fill-add me-2"></i>Form Tambah User</h5>
    </div>
    <div class="card-body">
      <form on:submit={handleSubmit}>
        <div class="mb-3">
          <label class="form-label" for="nama">Nama</label>
          <input id="nama" type="text" class="form-control" bind:value={nama} required />
        </div>

        <div class="mb-3">
          <label class="form-label" for="tanggalLahir">Tanggal Lahir</label>
          <input id="tanggalLahir" type="date" class="form-control" bind:value={tanggalLahir} required />
        </div>

        <div class="mb-3">
          <label class="form-label" for="jenisKelamin">Jenis Kelamin</label>
          <select id="jenisKelamin" class="form-select" bind:value={jenisKelamin} required>
            <option value="" disabled selected>Pilih Jenis Kelamin</option>
            <option value="Laki-laki">Laki-laki</option>
            <option value="Perempuan">Perempuan</option>
          </select>
        </div>

        <div class="mb-3">
          <label class="form-label" for="pekerjaan">Pekerjaan</label>
          <input id="pekerjaan" type="text" class="form-control" bind:value={pekerjaan} required />
        </div>

        <div class="mb-3">
          <label class="form-label" for="foto">Foto</label>
          <input id="foto" type="file" class="form-control" bind:files={foto} accept="image/*" />
        </div>

        <button type="submit" class="btn btn-primary">
          <i class="bi bi-save me-1"></i> Simpan
        </button>
        
        <a href="/" class="btn btn-secondary position-absolute end-0 bottom-0 mb-3 me-3">Kembali</a>
      </form>
    </div>
  </div>
</div>
