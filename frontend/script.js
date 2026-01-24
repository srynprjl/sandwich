let msg = document.getElementById('msg')

let cat = async () => {
  let cat_select = document.getElementById('category-select')
  const response = await fetch("http://localhost:5000/api/category/", {
    method: "GET"
  })
  if (!response.ok) {
    msg.innerText = "Error"
  }
  const result = await response.json();
  data = result.data
  data.forEach(element => {
    let id = element[0]
    let name = element[1]
    console.log(id, name)
    option = document.createElement('option')
    option.value = id
    option.innerText = name
    cat_select.appendChild(option)
  });
}
cat()

let cat_btn = document.getElementById('category-submit')
cat_btn.addEventListener('click', async (e) => {
  let cat_name = document.getElementById('category-name').value
  let cat_sh = document.getElementById('category-sh').value
  console.log(cat_name, cat_sh)
  e.preventDefault()
  if (cat_name != "" || cat_sh != "") {
    const response = await fetch("http://localhost:5000/api/category/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name: cat_name, shorthand: cat_sh })
    })
    if (!response.ok) {
      msg.innerText = "Error"
    }
    const result = await response.json();
    console.log(result)
    msg.innerText = result.message
  }
})

let project_btn = document.getElementById('project-btn')
project_btn.addEventListener('click', async (e) => {
  e.preventDefault()
  let cat_select = document.getElementById('category-select').value
  let name = document.getElementById('project-name').value
  let desc = document.getElementById('project-description').value
  let path = document.getElementById('project-path').value
  let fav = document.getElementById('project-fav').checked
  let com = document.getElementById('project-completed').checked

  let favo = fav ? 1 : 0
  let comp = com ? 1 : 0
  data = {
    name,
    description: desc,
    path,
    favorite: favo,
    completed: comp
  }
  if (data.name != "" || data.description != "" || path != "") {
    const response = await fetch(`http://localhost:5000/api/category/${cat_select}/project`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data)
    })
    if (!response.ok) {
      msg.innerText = "Error"
    }
    const result = await response.json();
    console.log(result)
    msg.innerText = `${result.message} || ${data.name}`
  }
})
