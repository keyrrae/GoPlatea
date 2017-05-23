/**
 * Created by xuanwang on 5/9/17.
 */
import React from 'react';
import { Editor, EditorState} from 'draft-js';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import axios from 'axios';

class CodeEditor extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      editorState: EditorState.createEmpty(),
      text: "",
      linenums: null,
      languageOption: "php",
      exeResult: [],
    };

    this.focus = () => this.refs.editor.focus();

    this.onChange = (editorState) => {
      this.setState({editorState});
      const content = this.state.editorState.getCurrentContent().getPlainText();
      //language=JSRegexp
      const numOfLine = (content.split("\n").length + 1);
      this.setState({linenums: this.genLineNums(numOfLine)});
    };

    this.logState = () => {
      this.setState({text:"log"});
    };

    this.runCode = () => {
      const req = {
        language: this.state.languageOption,
        code: this.state.editorState.getCurrentContent().getPlainText()
      };
      axios({
        method: 'post',
        url: 'http://localhost:8000',
        //url: 'http://ec2-54-215-223-77.us-west-1.compute.amazonaws.com:8000',
        data: JSON.stringify(req)
      })
          .then((response) => {
            this.setState({ exeResult: response.data });
            console.log(response)
          });
    };

    this.clearEditor = () => {
        this.setState({editorState: EditorState.createEmpty(), exeResult: []});
    };

    this.genLineNums = (numOfLine) => {
      let res = [];
      for(let i = 1; i <= numOfLine; i++){
        res.push(i);
      }
      return (
          <div>
            {res.map((i) => <div key={i}>{i}</div>)}
          </div>
      );
    };
  }

  render() {
    return (
        <div style={styles.root}>
            <div style={styles.container}>
                <div style={styles.linenum}>
                  {this.state.linenums}
                </div>
                <div style={styles.editor} onClick={this.focus}>
                    <Editor
                        editorState={this.state.editorState}
                        onChange={this.onChange}
                        ref="editor"
                    />
                </div>
            </div>
            <div>
                <select value={this.state.languageOption}
                        onChange={ event => this.setState({languageOption: event.target.value}) }>
                    <option value="php">PHP Zend and HHVM</option>
                    <option value="hack">HHVM</option>
                </select>
            </div>
            <input
                onClick={this.runCode}
                style={styles.button}
                type="button"
                value="Run Code"
            />
            <input
                onClick={this.clearEditor}
                style={styles.button}
                type="button"
                value="Clear"
            />
            <input
                onClick={this.logState}
                style={styles.button}
                type="button"
                value="About"
            />
            <div style={{paddingTop: 20}}>
              <ResultTabs content={this.state.exeResult}/>
            </div>
        </div>
    );
  }
}

class ResultTabs extends React.Component {
  constructor(props) {
    super(props);
    this.state = { tabIndex: 0 };
  }
  render() {
    if (this.props.content.length === 0){
      return (
          <div/>
      );
    }
    return (
        <Tabs
            selectedIndex={this.state.tabIndex}
            onSelect={
              tabIndex => this.setState({ tabIndex })
            }
        >
          <TabList>
            {this.props.content.map((res) =>
              <Tab key={res["name"]}>
                {res["name"]}
              </Tab>
            )}
          </TabList>
          {this.props.content.map((res) =>
              <TabPanel key={res["name"]}>
                <h4>Output</h4>
                <p>{res["output"]}</p>
                <h4>Execution Time</h4>
                <p>{'User:  ' + res["time"]["user"]+'s'}</p>
                <p>{'System:' + res["time"]["sys"]+'s'}</p>
              </TabPanel>
          )}
        </Tabs>
    );
  }
}

const styles = {
  root: {
    fontFamily: '\'Helvetica\', sans-serif',
    padding: 20,
    width: 800,
  },
  container: {
    display: 'flex',
    flexDirection: 'row',
  },
  linenum:{
    border: '1px solid #ccc',
    minHeight: 200,
    padding: 10,
    whiteSpaceTreatment: 'pre',
    fontSize: 15,
  },
  editor: {
    border: '1px solid #ccc',
    cursor: 'text',
    flex: 1,
    minHeight: 200,
    padding: 10,
    fontSize: 15,
  },
  button: {
    marginTop: 10,
    fontSize: 24,
    textAlign: 'center',
  },
  codeline: {
    whiteSpaceTreatment: 'nowrap'
  }
};

export default CodeEditor;