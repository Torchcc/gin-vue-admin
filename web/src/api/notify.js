import service from '@/utils/request'


// send sms to notify appointment ok
export const notifyAppointmentOk = (data) => {
  return service({
    url: "/notify/sms/appointment_ok/",
    method: 'post',
    data
  })
}
