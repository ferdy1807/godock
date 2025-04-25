<script lang="ts">
    export let users: {
      id: number;
      name: string;
      birthdate: string;
      gender: string;
      job: string;
      photo: string;
    }[] = [];  // Array of users
    
    import { goto } from '$app/navigation';
  
    function editUser(id: number) {
      goto(`/edit/${id}`);
    }
  </script>
  
  <h2 class="mb-3">Daftar Pengguna</h2>
  
  <div class="table-responsive">
    <table class="table table-bordered table-striped align-middle">
      <thead class="table-dark">
        <tr>
          <th>No</th>
          <th>Foto</th>
          <th>Nama</th>
          <th>Tanggal Lahir</th>
          <th>Jenis Kelamin</th>
          <th>Pekerjaan</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        {#each users as user, index}
          <tr>
            <td>{index + 1}</td>
            <td>
              {#if user.photo}
                <img src={"http://localhost:8080/assets/" + user.photo} width="60" class="rounded" alt="Foto" />
              {:else}
                <span>No Foto</span>
              {/if}
            </td>
            <td>{user.name}</td>
            <td>{user.birthdate}</td>
            <td>{user.gender}</td>
            <td>{user.job}</td>
            <td>
              <button class="btn btn-primary btn-sm" on:click={() => editUser(user.id)}>Edit</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  