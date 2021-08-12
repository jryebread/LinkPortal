import React, { Component } from 'react';
import LinkElement from './LinkElement.js'

class LinksArray extends Component {
    constructor(props) {
        super(props);

        this.state = {
            links: [],
            isLoaded: false
        }
    }

    componentDidMount() {
        const result  = fetchAsync("https://api.jamesr.me/users/james")
        result.then((result) => {   
            let links = []
            result.forEach(link => {
                console.log(link["Link"])
                links.push(link)
            })
            this.setState({ links, isLoaded:true});
        });
    }


   render() {
    const links = this.state.links;
    console.log(links, this.state.isLoaded)
    const linksComponents = links.map(link => <LinkElement key={link["ID"]} url={link["Link"]} />);

    return (
        <div>
            <ul>{linksComponents}</ul>
        </div>
      );
   }

}
    
async function fetchAsync (url) {
    let response = await fetch(url);
    let data = await response.json();
    return data;
}

function RetrieveAndListURLs() {
    fetchAsync("").then(links => {
        links.forEach(link => {
        console.log(link["Link"])
            this.state.links.push(<LinkElement url={link["Link"]}/>)
        })
    })
}
export default LinksArray;
