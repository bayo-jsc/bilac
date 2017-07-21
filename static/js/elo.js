new Vue({
  el: "#tf",
  delimiters: ['${', '}'],

  data: {
    members: [],
    color: ['#26C6DA', '#4DD0E1', '#80DEEA', '#B2EBF2']
  },

  mounted() {
    let loader = document.getElementById("preloader")
    loader.outerHTML = ""

    this.getMembers();
  },

  methods: {
    getMembers() {
      axios.get('api/v2/members',{
        params: {
          sort: "-elo"
        }
      })
        .then(res => {
          this.members = res.data
        }, err => {
          console.log(err)
        })
    }
  }
})
