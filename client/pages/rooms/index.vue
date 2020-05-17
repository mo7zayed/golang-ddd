<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <h3>All Rooms</h3>
        <hr>
        <nuxt-link to="/rooms/create" class="btn btn-primary">
          Create Room
        </nuxt-link>
        <hr>
        <table class="table table-bordered table-hover" v-if="rooms.length">
          <thead class="thead-dark">
            <tr>
              <th scope="col">#</th>
              <th scope="col">Room name</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(room, index) in rooms" :key="'room-' + room.id">
              <th scope="row">
                #{{ room.id }}
              </th>
              <td>
                {{ room.name }}
              </td>
              <td>
                <a href="#" class="btn btn-info">
                  Join
                </a>
                <a href="#" class="btn btn-danger" @click.prevent="deleteRoom(room.id, index)">
                  Delete
                </a>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="alert alert-danger" v-else>
          <h4>
            Looks like there is no rooms for chatting but the good news is
            <nuxt-link to="/rooms/create">
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
        rooms: [],
      }
    },
    async asyncData({
      params,
      app
    }) {
      try {
        let response = await app.$axios.get('groups')
        return {
          rooms: response.data.payload.data,
        }
      } catch (err) {
        console.error(err);
      }
    },
    methods: {
      async deleteRoom(room_id, index) {
        try {
          await this.$axios.delete('groups/' + room_id)
          this.rooms.splice(index, 1)
        } catch (err) {
          console.error(err);
        }
      }
    }
  }
</script>
