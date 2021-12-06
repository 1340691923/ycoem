export function getImageName(list,index){
  if (list.length == 0){
    return ""
  }
  return list[index].url
}
