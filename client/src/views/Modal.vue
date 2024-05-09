<script setup>
import { ref, computed} from 'vue'
import HeaderView from './Header.vue'
import FooterView from './Footer.vue'
import Swal from 'sweetalert2'
import { usePengunjungStore } from '../stores/pengunjung'

const pengunjungStore = usePengunjungStore()

const isOpen = ref(false);
const SelectedKategori = ref('Pilih Kategori');

const toggleDropdown = () => {
    isOpen.value = !isOpen.value;
};
const selectedKategoriOption = (kategori) => {
    SelectedKategori.value = kategori;
    isOpen.value = false;
};

const handler = () => {
    if (SelectedKategori.value === 'Pilih Kategori') {
        Swal.fire('Gagal!', 'Pilih kategori pengunjung', 'error')
        return
    } else if (SelectedKategori.value === 'Alumni') {
        form.value.asal= 'SMK Negeri 7 Yogyakarta'
    } else if (SelectedKategori.value === 'Umum') {
        form.value.jurusan = 'Umum'
    }
}
const validasiNoHP = () => {
    const phoneRegex = /^[0-9]{10,13}$/;
    return phoneRegex.test(form.value.no_hp)
}

const pengunjung = ref([])
const errorMessage = ref('')
const form = computed(() => {
    return {
        nama: '',
        no_hp: '',
        kategori: SelectedKategori.value,
        tahun_lulus: '',
        jurusan: '',
        asal: ''
    };
});

const Register = async () => {
    if (!form.value.nama || !form.value.no_hp || !form.value.tahun_lulus) {
        Swal.fire('Gagal!', 'Data tidak boleh kosong', 'error')
        return
    } else if (validasiNoHP()) {
        try {
            let response = await pengunjungStore.Register(form.value)
            errorMessage.value = response.message
            if (errorMessage.value === 'Success') {
                Swal.fire({
                    icon: 'success',
                    title: 'TerimaKasih!',
                    text: '',
                    showConfirmButton: false,
                    timer: 1500
                })
                form.value = ''
                SelectedKategori.value = 'Pilih Kategori'
            } else {
                Swal.fire({
                    icon: 'error',
                    title: 'Error!',
                    text: errorMessage.value,
                    showConfirmButton: false,
                    timer: 1500
                })
            }
        } catch (error) {
            alert(error.message)
        }
    } else if (!validasiNoHP()) {
        Swal.fire({
                    icon: 'error',
                    title: 'Error!',
                    text: 'No HP/WA tidak valid',
                    showConfirmButton: false,
                    timer: 1500
                })
    }
}
</script>
<template>
    <main class="flex flex-col min-h-screen">
    <header>
      <HeaderView />
    </header>            
    <div class="flex-grow w-full flex items-center justify-center px-5 mx-auto bg-[url('../assets/image/bg.png')]">
        <div class="w-screen overflow-hidden">
            <div class="flex items-center justify-center py-4 text-xl font-semibold">
                <svg class="mr-2 h-5 w-5" fill="#000000" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 974.505 974.505" xml:space="preserve"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <g> <path d="M284.22,974.505h624.5c33.3,0,60.2-27,60.2-60.2v-526c0-24.9-20.101-45-45-45h-260.5c-24.9,0-45-20.1-45-45v-253.2 c0-24.9-20.101-45-45-45h-289.2c-33.3,0-60.2,27-60.2,60.2v77.2l404.101,401.3c10.6,10.601,18.699,23.7,23.399,37.9l47.7,145.7 c4,12.199,4.9,25.199,2.7,37.8c-2,11.6-6.8,22.899-13.8,32.5c-14.801,20.399-38.5,32.6-63.5,32.6c-6.4,0-12.9-0.8-19.101-2.399 l-154.899-39c-16.601-4.2-31.9-12.801-44-24.801l-182.7-181.399v336.5C223.92,947.505,250.92,974.505,284.22,974.505z"></path> <path d="M949.02,274.705c17.8,0,26.7-21.5,14.101-34.1L728.52,5.905c-12.6-12.6-34.1-3.7-34.1,14.1v234.6c0,11.1,9,20.1,20.1,20.1 H949.02z"></path> <path d="M120.82,140.105c-12.7,0-25.5,4.2-33.9,12.6l-68.8,68.8c-16.6,16.6-16.9,50.3-0.6,66.6l60,59.6l136-136l-60-59.6 C145.42,144.105,133.22,140.105,120.82,140.105z"></path> <polygon points="417.82,685.705 553.82,549.705 256.12,254.005 223.92,286.205 120.12,390.005 223.92,493.205 "></polygon> <path d="M620.22,764.805c1.5,0.4,3,0.601,4.5,0.601c12,0,21.5-12,17.6-24.2l-47.8-145.7c-0.1-0.4-0.3-0.8-0.399-1.1l-130.9,130.8 c0.7,0.2,1.4,0.399,2,0.6L620.22,764.805z"></path> </g> </g></svg>
                <h1>Presensi</h1>
            </div>
            <div class="flex items-center justify-center">
                <div class="max-w-md border bg-white border-gray-200 rounded flex-1 z-20 p-7">
            <div class="mb-2">
                <h1 class="font-bold text-lg">Data Diri</h1>
                <hr>
                <div class="pt-1" @click="toggleDropdown()">
                    <span class="text-xs">Kategori Pengunjung</span>
                    <button
                        class="py-1.5 px-2 bg-gray-400 text-white rounded hover:shadow flex items-center justify-between w-32 mb-2 text-sm">
                        <span>{{ SelectedKategori }}</span>
                        <svg class="h-4 w-4" :class="{ 'transform rotate-180': isOpen }" fill="currentColor"
                            viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                            <path fill-rule="evenodd" d="M10 15l5-5H5l5 5z" clip-rule="evenodd"></path>
                        </svg>
                    </button>
                    <div v-if="isOpen"
                        class="dropdown-menu absolute z-10 mt-1 max-w-sm bg-white border border-gray-300 rounded-md shadow-lg text-sm">
                        <div v-for="(kategori, index) in ['Alumni', 'Umum']" :key="index"
                            @click="selectedKategoriOption(kategori); toggleDropdown();handler()"
                            class="cursor-pointer px-4 py-2 w-32 hover:bg-gray-100 border-b border-gray-300">
                            {{ kategori }}
                        </div>
                    </div>
                </div>
                <div :hidden="SelectedKategori === 'Pilih Kategori'">
                    <div class="py-1">
                        <p class="text-sm">Nama Lengkap</p>
                        <input v-model="form.nama" class="w-full px-1.5 py-1 border border-gray-400 rounded text-sm" @input="form.nama=$event.target.value.charAt(0).toUpperCase() + $event.target.value.slice(1)" type="text">
                    </div>
                    <div class="py-1">
                        <p class="text-sm">No HP/WA</p>
                        <input v-model="form.no_hp" class="w-full px-1.5 py-1 border border-gray-400 rounded text-sm [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none" 
                        type="text" inputmode="numeric" pattern="[0-9]*">
                    </div>
                    <div class="py-1">
                        <p class="text-sm">Tahun Lulus</p>
                        <input v-model="form.tahun_lulus" class="w-full px-1.5 py-1 border border-gray-400 rounded text-sm [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none" 
                            type="text" inputmode="numeric" pattern="[0-9]">
                    </div>
                    <div :hidden="SelectedKategori != 'Alumni'" class="py-1">
                        <p class="text-sm">Jurusan</p>
                        <input v-model="form.jurusan" class="w-full px-1.5 py-1 border border-gray-400 rounded text-sm" @input="form.jurusan=$event.target.value.toUpperCase()" type="text">
                    </div>
                    <div :hidden="SelectedKategori != 'Umum'" class="py-1">
                        <p class="text-sm">Asal Sekolah/Instansi</p>
                        <input v-model="form.asal" class="w-full px-1.5 py-1 border border-gray-400 rounded text-sm" @input="form.asal=$event.target.value.toUpperCase()" type="text">
                    </div>
                </div>
            </div>
            <button @click="Register"
                            class="bg-blue-500 text-white px-3 py-1.5 rounded-md text-sm">Kirim</button>
        </div>
            </div>
        </div>
    </div>
    <footer class="bg-gray-900">
      <FooterView />
    </footer>
  </main>
</template>