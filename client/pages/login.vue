<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-md-8">
        <div class="card">
          <div class="card-header">Login</div>

          <div class="card-body">
            <ValidationErrors :errors="errors" />

            <form @submit.prevent="submit()">
              <div class="form-group row">
                <label for="email" class="col-md-4 col-form-label text-md-right">E-Mail Address</label>

                <div class="col-md-6">
                  <input id="email" type="email" v-model="form.email" class="form-control" name="email" required autocomplete="email"
                    autofocus>
                </div>
              </div>

              <div class="form-group row">
                <label for="password" class="col-md-4 col-form-label text-md-right">
                  Password
                </label>

                <div class="col-md-6">
                  <input id="password" type="password" v-model="form.password" class="form-control" name="password" required
                    autocomplete="current-password">
                </div>
              </div>

              <div class="form-group row mb-0">
                <div class="col-md-8 offset-md-4">
                  <button type="submit" class="btn btn-primary">
                    Login
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
          email: '',
          password: '',
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
          $data.append("email", this.form.email)
          $data.append("password", this.form.password)

          await this.$auth.loginWith('local', {
            data: $data
          })
        } catch (err) {
          this.errors = err.response.data.errors;
        }
      },
    },
  }
</script>
