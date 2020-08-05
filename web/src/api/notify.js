import service from '@/utils/request'


// send sms to notify appointment ok
export const notifyAppointmentOk = (data) => {
  return service({
    url: "/notify/sms/appointment_ok/",
    method: 'post',
    data
  })
}

// send sms to notify refund ok
export const notifyRefundOk = (data) => {
  return service({
    url: "/notify/sms/refund_ok/",
    method: 'post',
    data
  })
}

