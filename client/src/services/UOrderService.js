import Api from '@/services/Api.js'

export default {
  searchMed (userdata) {
    console.log(userdata)
    return Api.post('/search', userdata)
  }
}
