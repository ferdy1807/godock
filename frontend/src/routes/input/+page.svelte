<script lang="ts">
  import UserForm from '../../components/UserForm.svelte';
  import { goto } from '$app/navigation';

  // Data default untuk user
  let user = {
    id: 0,
    name: '',
    birthdate: '',
    gender: 'Laki-laki',
    job: '',
    photo: ''
  };

  // Fungsi untuk mengirim data pengguna ke API
  async function submitUser(formData: FormData) {
    const res = await fetch('http://localhost:8080/users', {
      method: 'POST',
      body: formData
    });

    if (res.ok) {
      goto('/');  // Kembali ke halaman utama setelah berhasil
    } else {
      console.error('Gagal menambah pengguna');
    }
  }
</script>

<div class="container mt-5">
  <h1>Tambah User Baru</h1>
  <div class="card p-4">
    <UserForm {user} onSubmit={submitUser} />  <!-- Mengirimkan user ke UserForm -->
  </div>
</div>
