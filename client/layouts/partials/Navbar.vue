<template>
  <nav class="navbar navbar-expand-md navbar-dark bg-dark shadow-sm">
    <div class="container">
      <a class="navbar-brand" href="">
        Chat App
      </a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <!-- Left Side Of Navbar -->
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <nuxt-link class="nav-link" to="/">
              Home
            </nuxt-link>
          </li>

          <template v-if="$auth.loggedIn">
            <li class="nav-item">
              <nuxt-link class="nav-link" to="/rooms/">
                Rooms
              </nuxt-link>
            </li>

            <li class="nav-item" v-if="$auth.user.user_type == 'admin'">
              <nuxt-link class="nav-link" to="/users/">
                Users
              </nuxt-link>
            </li>
          </template>
        </ul>

        <!-- Right Side Of Navbar -->
        <ul class="navbar-nav ml-auto" v-if="$auth.loggedIn">
          <b-nav-item-dropdown :text="$auth.user.name" right>
            <b-dropdown-item href="#" @click.prevent="logout()">Logout</b-dropdown-item>
          </b-nav-item-dropdown>
        </ul>
        <ul class="navbar-nav ml-auto" v-else>
          <li class="nav-item">
            <nuxt-link class="nav-link" to="/login">
              Login
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link class="nav-link" to="/register">
              Register
            </nuxt-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
  export default {
    methods: {
      logout() {
        this.$auth.logout()
      }
    },
  }

</script>
