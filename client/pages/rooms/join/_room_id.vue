<template>
  <div class="container">
    <div class="chat_area">
      <ul class="list-unstyled" v-if="messages.length">
        <li v-for="(m, index) in messages" :key="'line-' + index">
          <div class="chat-body1">
            <p>
              <b>{{ m.username }}:</b>
              {{ m.message }}
            </p>
          </div>
        </li>
      </ul>
      <div class="alert alert-info" v-else>
        No Messages In This Room Yet. Be The First
      </div>
    </div>

    <div class="message_write">
      <form @submit.prevent="send()">
        <textarea class="form-control" v-model="form.message" placeholder="type a message"></textarea>
        <div class="chat_bottom">
          <button type="submit" class="pull-right btn btn-success">
            Send
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        form: {
          message: '',
        },
        ws: null,
        wsURL: 'ws://127.0.0.1:5000/api/v1/ws',
        room_id: null,
        messages: [],
        errors: {}
      }
    },
    async asyncData({
      params,
      app
    }) {
      try {
        let room_id = params.room_id

        let response = await app.$axios.get('messages-by-channel/' + `channel.room.${room_id}`)

        return {
          messages: response.data.payload.data,
          room_id: params.room_id,
        }
      } catch (err) {
        console.error(err);
      }
    },
    mounted() {
      this.ws = new WebSocket(this.wsURL);

      this.ws.onopen = function () {
        console.log('Connected')
      }

      this.ws.onmessage = (e) => {
        let $data = JSON.parse(e.data);
        if ($data.channel == this.getChannelName()) {
          this.messages.push($data)
        }
      }
    },
    methods: {
      send() {
        if (this.form.message != '') {
          this.ws.send(JSON.stringify(this.createMessage(this.form.message)));

          this.form.message = ''; // Reset newMsg
        }
      },
      getChannelName() {
        return `channel.room.${this.room_id}`
      },
      createMessage(message) {
        return {
          username: this.$auth.user.name,
          user_id: this.$auth.user.id,
          channel: this.getChannelName(),
          message: message,
        }
      },
    },
  }

</script>

<style scoped>
  .all_conversation button {
    background: #f5f3f3 none repeat scroll 0 0;
    border: 1px solid #dddddd;
    height: 38px;
    text-align: left;
    width: 100%;
  }

  .all_conversation i {
    background: #e9e7e8 none repeat scroll 0 0;
    border-radius: 100px;
    color: #636363;
    font-size: 17px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    width: 30px;
  }

  .all_conversation .caret {
    bottom: 0;
    margin: auto;
    position: absolute;
    right: 15px;
    top: 0;
  }

  .all_conversation .dropdown-menu {
    background: #f5f3f3 none repeat scroll 0 0;
    border-radius: 0;
    margin-top: 0;
    padding: 0;
    width: 100%;
  }

  .all_conversation ul li {
    border-bottom: 1px solid #dddddd;
    line-height: normal;
    width: 100%;
  }

  .all_conversation ul li a:hover {
    background: #dddddd none repeat scroll 0 0;
    color: #333;
  }

  .all_conversation ul li a {
    color: #333;
    line-height: 30px;
    padding: 3px 20px;
  }

  .top_nav {
    overflow: visible;
  }

  .chat-img img {
    height: 34px;
    width: 34px;
  }

  .sub_menu_ {
    background: #e8e6e7 none repeat scroll 0 0;
    left: 100%;
    max-width: 233px;
    position: absolute;
    width: 100%;
  }

  .sub_menu_ {
    background: #f5f3f3 none repeat scroll 0 0;
    border: 1px solid rgba(0, 0, 0, 0.15);
    display: none;
    left: 100%;
    margin-left: 0;
    max-width: 233px;
    position: absolute;
    top: 0;
    width: 100%;
  }

  .all_conversation ul li:hover .sub_menu_ {
    display: block;
  }

  .new_message_head button {
    background: rgba(0, 0, 0, 0) none repeat scroll 0 0;
    border: medium none;
  }

  .new_message_head {
    background: #f5f3f3 none repeat scroll 0 0;
    float: left;
    font-size: 13px;
    font-weight: 600;
    padding: 18px 10px;
    width: 100%;
  }

  .message_section {
    border: 1px solid #dddddd;
  }

  .chat_area {
    float: left;
    height: 600px;
    overflow-x: hidden;
    overflow-y: auto;
    width: 100%;
  }

  .chat_area li {
    padding: 14px 14px 0;
  }

  .chat_area li .chat-img1 img {
    height: 40px;
    width: 40px;
  }

  .chat-body1 p {
    background: #fbf9fa none repeat scroll 0 0;
    padding: 10px;
  }

  .chat_area .admin_chat .chat-body1 {
    margin-left: 0;
    margin-right: 50px;
  }

  .chat_area li:last-child {
    padding-bottom: 10px;
  }

  .message_write {
    background: #f5f3f3 none repeat scroll 0 0;
    float: left;
    padding: 15px;
    width: 100%;
  }

  .message_write textarea.form-control {
    height: 70px;
    padding: 10px;
  }

  .chat_bottom {
    float: left;
    margin-top: 13px;
    width: 100%;
  }

  .upload_btn {
    color: #777777;
  }

  .sub_menu_>li a,
  .sub_menu_>li {
    float: left;
    width: 100%;
  }

</style>
