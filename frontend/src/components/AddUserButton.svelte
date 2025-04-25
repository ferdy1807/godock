<script>
    import { goto } from '$app/navigation';
  
    function handleClick() {
      goto('/input');
    }
  </script>
  
  <button class="btn btn-success mb-4" on:click={handleClick}>Tambah User</button>
  