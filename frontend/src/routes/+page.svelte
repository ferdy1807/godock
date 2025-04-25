<script lang="ts">
    import { onMount } from 'svelte';
    import AddUserButton from '../components/AddUserButton.svelte';
    import UserList from '../components/UserList.svelte';
  
    // Mendefinisikan interface untuk user
    interface User {
      id: number;
      name: string;
      birthdate: string;
      gender: string;
      job: string;
      photo: string;
    }
  
    // Mendeklarasikan users dengan tipe User[]
    let users: User[] = []; // Array of User objects
  
    onMount(async () => {
      const res = await fetch('http://localhost:8080/users');
      if (res.ok) {
        users = await res.json();
      } else {
        console.error('Gagal memuat data pengguna');
      }
    });
  </script>
  
  <h1>Dashboard</h1>
  <AddUserButton />
  <hr />
  <UserList {users} />
  