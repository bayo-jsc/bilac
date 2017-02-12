new Vue({
  el: '#tf',
  delimiters: ['${', '}'],

  data: {
    members: [],
    newMem: {},
  },

  mounted() {
    axios.get('api/v1/members')
      .then(res => {
        this.members = res.data
      }, err => {
        console.log(err)
      })
  },

  methods: {
    createMember() {
      axios.post('/api/v1/members', {
        username: this.newMem.username
      }).then(res => {
          this.members.push(res.data)
          this.newMem = {}
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
