<script lang="ts">
  import UserForm from '../../../components/UserForm.svelte';  // Mengimpor UserForm

  // Mendapatkan data pengguna melalui load()
  export async function load({ params }: { params: { id: string } }) {
    const { id } = params;
    const res = await fetch(`http://localhost:8080/users/${id}`);

    if (!res.ok) {
      return {
        status: res.status,
        error: new Error('Data pengguna tidak ditemukan')
      };
    }

    const user = await res.json();
    return { user };
  }

  export let user: { id: number, name: string, birthdate: string, gender: string, job: string, photo: string };

  // Fungsi untuk memperbarui data pengguna
  async function updateUser(formData: FormData) {
    const res = await fetch(`http://localhost:8080/users/${user.id}`, {
      method: 'PUT',
      body: formData
    });

    if (!res.ok) {
      alert('Gagal memperbarui pengguna');
    } else {
      alert('Pengguna berhasil diperbarui');
    }
  }
</script>

{#if user}
  <div class="container mt-5">
    <h1>Edit Pengguna</h1>
    <div class="card p-4">
      <!-- Mengoper user dan onSubmit ke komponen UserForm -->
      <UserForm user={user} onSubmit={updateUser} />
    </div>
  </div>
{:else}
  <p>Data pengguna tidak ditemukan.</p>
{/if}
