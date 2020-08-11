import service from '@/utils/request'

export const listRegions = (params) => {
  return service({
    url: "/region/getRegionList",
    method: 'get',
    params
  })
}

