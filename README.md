# Fortune cookie 🍪

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

Simple react app for learning purposes that gives you random fortunes

# Setting up

This project utilizes a local instance of [fortune-api](https://github.com/sarah256/fortune-api) to get the fortunes

## Installing dependencies:

### `📁 Fortune cookie`

```bash
$ npm install
```

### `📁 Fortune API local instance`

### 1️⃣ Install [GO language 1.7+](https://go.dev/dl/)

For windows machines, download the `.msi`

### 2️⃣ Install fortune-api dependencies

```bash
$ go get github.com/gorilla/mux
```

#

### To run the project as-is, you will need to install globally cross-env

```bash
$ npm install -g cross-env@7.0.3
```

Or just remove the `cross-env BROWSER=none` from the start script in `package.json`

# Running the app

### `📁 Fortune API local instance`

```bash
$ go run main.go
```

### `📁 Fortune cookie`

```bash
$ npm start
```

### Obs: The app translate all the fortunes to `brazilian portuguese`

### if you want it to display the fortunes in english, change the lines from [clicarBotao](https://github.com/JuniorLoch/react-fortune-cookie/blob/9b1836adacb2296c3d075c7a6324ccd4ec0b7e4a/Fortune%20cookie/src/App.js#L34)



### from:

```javascript
msg: await translate(await this.msgAleatoria(), 'Portuguese'),
```

### to:

```javascript
msg: await this.msgAleatoria(),
```
