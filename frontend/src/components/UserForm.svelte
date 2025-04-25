<script lang="ts">
    import { goto } from '$app/navigation'; // Untuk navigasi
    export let user: {
      id: number | null;
      name: string;
      birthdate: string;
      gender: string;
      job: string;
      photo: string;
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
            }, 2000);
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
  
  <!-- Container untuk form -->
  <div class="container mt-4">
    <h2>Edit User Information</h2>
    <form class="position-relative" on:submit={handleSubmit}>
      <!-- Field Name -->
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input type="text" id="name" class="form-control" bind:value={user.name} />
      </div>
  
      <!-- Field Birthdate -->
      <div class="mb-3">
        <label for="birthdate" class="form-label">Birthdate</label>
        <input type="date" id="birthdate" class="form-control" bind:value={user.birthdate} />
      </div>
  
      <!-- Field Gender -->
      <div class="mb-3">
        <label for="gender" class="form-label">Gender</label>
        <input type="text" id="gender" class="form-control" bind:value={user.gender} />
      </div>
  
      <!-- Field Job -->
      <div class="mb-3">
        <label for="job" class="form-label">Job</label>
        <input type="text" id="job" class="form-control" bind:value={user.job} />
      </div>
  
      <!-- Field Photo -->
      <div class="mb-3">
        <label for="photo" class="form-label">Photo URL</label>
        <input type="text" id="photo" class="form-control" bind:value={user.photo} />
      </div>
  
      <!-- Tombol Save -->
      <button type="submit" class="btn btn-primary">Save</button>
  
      <!-- Tombol Kembali ke Homepage -->
      <a href="/" class="btn btn-secondary position-absolute end-0 bottom-0 mb-3 me-3">Kembali</a>
    </form>
  </div>
  