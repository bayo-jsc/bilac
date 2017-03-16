new Vue({
  el: '#app',
  delimiters: ['${', '}'],

  data() {
    return {
      members: [],
      players: [],
      oldPlayers: [],
      teams: [],
      newMem: {},
      newMember: '',
      isViewScoreboard: true,
    }
  },

  mounted() {
    this.getMembers()
  },

  methods: {
    createMember() {
      axios.post('/api/v1/members', {
        username: this.newMember
      }).then(res => {
          this.members.push(res.data)
          this.$set(this, 'newMember', '')
        }, err => {
          console.log(err)
        })
    },

    destroyMember(index) {
      axios.delete('/api/v1/members/' + this.members[index].id)
        .then(res => {
          this.members.splice(index, 1)
        }, err => {
          console.log(err)
        })
    },

    getMembers() {
      axios.get('api/v1/members')
        .then(res => {
          this.$set(this, 'members', res.data)
        }, err => {
          console.log(err)
        })
    },

    draw() {
      axios.patch('/api/v1/draw')
        .then(res => {
          this.members = res.data
        }, err => {
          console.log(err)
        })
    },
  }
})
