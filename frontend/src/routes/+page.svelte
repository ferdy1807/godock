<script lang="ts">
  import { onMount } from 'svelte';
  import AddUserButton from '../components/AddUserButton.svelte';
  import UserList from '../components/UserList.svelte';

  interface User {
    id: number;
    name: string;
    birthdate: string;
    gender: string;
    job: string;
    photo: string;
  }

  let users: User[] = [];
  let userCount: number = 0;

  onMount(async () => {
    const userRes = await fetch('http://localhost:8080/users');
    if (userRes.ok) {
      users = await userRes.json();
    } else {
      console.error('Gagal memuat data pengguna');
    }

    const countRes = await fetch('http://localhost:8080/count');
    if (countRes.ok) {
      const data = await countRes.json();
      userCount = data.user_count;
    } else {
      console.error('Gagal memuat jumlah data');
    }
  });
</script>

<h1>Dashboard</h1>

<div class="d-flex justify-content-between mb-4">
  <!-- <h2>Total Users: {userCount}</h2> -->
  <AddUserButton />
</div>

<hr />
<UserList {users} />
