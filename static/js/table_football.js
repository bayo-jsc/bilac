new Vue({
  el: '#tf',
  delimiters: ['${', '}'],

  data: {
    members: [],
  },

  mounted() {
    axios.get('api/v1/members')
      .then(res => {
        this.members = res.data
      }, err => {
        console.log(err)
      })
  },
})
