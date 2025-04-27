<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores'; // Untuk mengakses parameter URL
  import { goto } from '$app/navigation'; // Untuk navigasi
  import UserForm from '../../../components/UserForm.svelte'; // Import komponen UserForm

  let user = {
    id: null,
    name: '',
    birthdate: '',
    gender: '',
    job: '',
    photo: ''
  };

  // Menangani pemuatan data berdasarkan parameter ID
  onMount(() => {
    const id = $page.params.id; // Mengambil parameter ID dari URL
    if (id) {
      fetchUser(Number(id)); // Memanggil fungsi fetchUser dengan ID sebagai angka
    }
  });

  // Fungsi untuk mengambil data user berdasarkan ID
  const fetchUser = async (userId: number) => {
    const res = await fetch(`http://localhost:8080/users/${userId}`);
    if (res.ok) {
      user = await res.json(); // Menyimpan data user ke dalam variabel user
    } else {
      console.error('Gagal mengambil data user');
    }
  };

  // Fungsi untuk menangani pengiriman form
  const handleSubmit = async (event: Event) => {
    event.preventDefault(); // Mencegah pengiriman form secara default

    // Periksa apakah user.id sudah ada, berarti update
    if (user.id) {
      try {
        const res = await fetch(`http://localhost:8080/users/${user.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(user), // Mengirimkan data user yang sudah diubah
        });

        if (res.ok) {
          alert('User updated successfully');
          setTimeout(() => {
            goto('/'); // Arahkan kembali ke halaman utama setelah alert
          }, 1000);
        } else {
          alert('Failed to update user');
        }
      } catch (error) {
        console.error('Error updating user:', error);
        alert('Error updating user');
      }
    }
  };
</script>

<!-- Gunakan komponen UserForm untuk menampilkan data -->
<UserForm {user} on:submit={handleSubmit} />
