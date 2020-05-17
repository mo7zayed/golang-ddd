<template>
  <div>
    <Navbar />
    <main class="py-4">
      <nuxt />
    </main>
  </div>
</template>

<script>
  import Navbar from '@/layouts/partials/Navbar'

  export default {
    components: {
      Navbar,
    },
  }
</script>
