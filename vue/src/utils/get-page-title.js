//获取title
import defaultSettings from '@/settings'

const title = defaultSettings.title || '烈焰鸟-单机'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
