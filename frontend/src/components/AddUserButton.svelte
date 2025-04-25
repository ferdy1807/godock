<script>
    import { goto } from '$app/navigation';
  
    function handleClick() {
      goto('/input');
    }
  </script>
  
  <button class="btn btn-success d-flex align-items-center gap-2 mb-4 shadow" on:click={handleClick}>
    <i class="bi bi-person-plus-fill"></i> Tambah User
  </button>
  