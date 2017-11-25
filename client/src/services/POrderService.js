import Api from '@/services/Api.js'

export default {
  getOrder (storeId) {
    return Api.post('/store/orders', storeId)
  } 
}
