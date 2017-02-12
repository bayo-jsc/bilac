new Vue({
  el: '#tf',
  delimiters: ['${', '}'],

  data: {
    members: [],
    newMem: {},
    errors: {},
  },

  mounted() {
    axios.get('api/v1/members')
      .then(res => {
        this.members = res.data
      }, err => {
        this.errors = err.responseJSON.error
      })
  },

  methods: {
    createMember() {
      axios.post('/api/v1/members', {
        username: this.newMem.username
      }).then(res => {
          this.errors = {}
          this.members.push(res.data)
        }, err => {
          this.errors = err.responseJSON.errors
        })
    }
  }
})
