<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-md-8">
        <div class="card">
          <div class="card-header">Create room</div>

          <div class="card-body">
            <ValidationErrors :errors="errors" />

            <form @submit.prevent="submit()">
              <div class="form-group row">
                <label for="name" class="col-md-4 col-form-label text-md-right">Name</label>

                <div class="col-md-6">
                  <input id="name" type="name" v-model="form.name" class="form-control" name="name" required autocomplete="name"
                    autofocus>
                </div>
              </div>

              <div class="form-group row mb-0">
                <div class="col-md-8 offset-md-4">
                  <button type="submit" class="btn btn-primary">
                    Create a room
                  </button>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import ValidationErrors from '@/components/global/ValidationErrors'

  export default {
    data() {
      return {
        form: {
          name: '',
        },
        errors: {}
      }
    },
    components: {
      ValidationErrors
    },
    methods: {
      async submit() {
        try {

          let $data = new FormData()
          $data.append("name", this.form.name)

          let response = await this.$axios.post('groups', $data)

          this.$router.push("/rooms")
        } catch (err) {
          this.errors = err.response.data.errors;
        }
      },
    },
  }
</script>
