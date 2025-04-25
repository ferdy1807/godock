<script lang="ts">
    export let user: {
      id: number;
      name: string;
      birthdate: string;
      gender: string;
      job: string;
      photo: string;
    };
    
    export let onSubmit: (formData: FormData) => void;
    
    let file: File | null = null;
  
    function handleSubmit() {
      const formData = new FormData();
      formData.append('name', user.name);
      formData.append('birthdate', user.birthdate);
      formData.append('gender', user.gender);
      formData.append('job', user.job);
      
      if (file) {
        formData.append('photo', file);
      }
    
      onSubmit(formData);
    }
  
    function handleFileChange(event: Event) {
      const input = event.target as HTMLInputElement;
      if (input?.files?.length) {
        file = input.files[0];
      }
    }
  </script>
  
  <form on:submit|preventDefault={handleSubmit}>
    <div class="mb-3">
      <label for="name" class="form-label">Nama</label>
      <input
        type="text"
        id="name"
        class="form-control"
        bind:value={user.name}
        placeholder="Masukkan nama"
        required
      />
    </div>
  
    <div class="mb-3">
      <label for="birthdate" class="form-label">Tanggal Lahir</label>
      <input
        type="date"
        id="birthdate"
        class="form-control"
        bind:value={user.birthdate}
        required
      />
    </div>
  
    <div class="mb-3">
      <label for="gender" class="form-label">Jenis Kelamin</label>
      <select
        id="gender"
        class="form-select"
        bind:value={user.gender}
        required
      >
        <option value="Laki-laki">Laki-laki</option>
        <option value="Perempuan">Perempuan</option>
      </select>
    </div>
  
    <div class="mb-3">
      <label for="job" class="form-label">Pekerjaan</label>
      <input
        type="text"
        id="job"
        class="form-control"
        bind:value={user.job}
        placeholder="Masukkan pekerjaan"
        required
      />
    </div>
  
    <div class="mb-3">
      <label for="photo" class="form-label">Foto</label>
      <input
        type="file"
        id="photo"
        class="form-control"
        on:change={handleFileChange}
      />
    </div>
  
    <button type="submit" class="btn btn-primary">Simpan</button>
  </form>
  