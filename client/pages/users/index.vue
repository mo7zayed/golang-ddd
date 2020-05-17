<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <h3>All Users</h3>
        <hr>
        <nuxt-link to="/users/create" class="btn btn-primary">
          Create User
        </nuxt-link>
        <hr>
        <table class="table table-bordered table-hover" v-if="users.length">
          <thead class="thead-dark">
            <tr>
              <th scope="col">#</th>
              <th scope="col">Name</th>
              <th scope="col">Email</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(user, index) in users" :key="'user-' + user.id">
              <th scope="row">
                #{{ user.id }}
              </th>
              <td>
                {{ user.name }}
              </td>
              <td>
                {{ user.email }}
              </td>
              <td>
                <a href="#" class="btn btn-danger" @click.prevent="deleteUser(user.id, index)">
                  Delete
                </a>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="alert alert-danger" v-else>
          <h4>
            Looks like there is no users
            <nuxt-link to="/users/create">
              You can create a one
            </nuxt-link>
          </h4>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        users: [],
      }
    },
    async asyncData({
      params,
      app
    }) {
      try {
        let response = await app.$axios.get('users')
        return {
          users: response.data.payload.data,
        }
      } catch (err) {
        console.error(err);
      }
    },
    methods: {
      async deleteUser(user_id, index) {
        try {
          await this.$axios.delete('users/' + user_id)
          this.users.splice(index, 1)
        } catch (err) {
          console.error(err);
        }
      }
    }
  }
</script>
