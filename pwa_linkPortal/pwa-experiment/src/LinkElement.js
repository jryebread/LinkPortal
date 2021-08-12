import React, { Component } from 'react';
import { LinkPreview } from '@dhaiwat10/react-link-preview';

function isValidHttpUrl(string) {
    let url;
    console.log(string)
    try {
      url = new URL(string);
    } catch (_) {
      return false;  
    }
  
    return url.protocol === "http:" || url.protocol === "https:";
}

function UrlExists(url) {
    var http = new XMLHttpRequest();
    http.open('HEAD', url, false);
    http.send();
    if (http.status == 200)
        return true
    else
        return false
}

class LinkElement extends Component {
    constructor(props) {
        super(props);

        this.state= {
            toRender: true
        }
        if (!isValidHttpUrl(this.props.url)) {
            console.log("Fake url!")
            this.setState({toRender: false})
        }
    }
   render() {
    const renderLink = this.state.toRender ? 
    <LinkPreview url={this.props.url} width='280px'/> 
                                : this.props.url
    return (
        <div>
{renderLink}
        </div>
    ); 
      
   }

}
export default LinkElement;
