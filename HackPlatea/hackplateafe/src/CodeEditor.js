/**
 * Created by xuanwang on 5/9/17.
 */
import React from 'react';
import { Editor, EditorState, getDefaultKeyBinding, KeyBindingUtil} from 'draft-js';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import axios from 'axios';

const {hasCommandModifier} = KeyBindingUtil;

function myKeyBindingFn(e) {
  if (e.keyCode === 13 /* Tab key */ && hasCommandModifier(e)) {
    return 'cmd_enter';
  }
  return getDefaultKeyBinding(e);
}


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

    this.handleKeyCommand = this.handleKeyCommand.bind(this);

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
        //url: 'http://localhost:8000',
        url: 'http://ec2-54-215-223-77.us-west-1.compute.amazonaws.com:8000',
        data: JSON.stringify(req)
      })
          .then((response) => {
            this.setState({ exeResult: response.data });
            console.log(response)
          });
    };

    this.clearEditor = () => {
        this.setState({
          editorState: EditorState.createEmpty(),
          exeResult: [],
          linenums: null
        });
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

  handleKeyCommand(command) {
    if (command === 'cmd_enter') {
      console.log('pressed');
      // Perform a state change to key press
      this.runCode();
      return 'handled';
    }
    return 'not-handled';
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
                        handleKeyCommand={this.handleKeyCommand}
                        keyBindingFn={myKeyBindingFn}
                        ref="editor"
                    />
                </div>
            </div>
            <div>
                <select value={this.state.languageOption}
                        onChange={ event => this.setState({languageOption: event.target.value})}
                        style={styles.select}
                >
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
    width: 600,
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
  select: {
    marginTop: 10,
    fontSize: 15,
    borderRadius: 4,
    backgroundColor: '#e2e2e2',
  },
  button: {
    marginTop: 10,
    marginRight: 10,
    borderRadius: 4,
    border: 'none',
    fontSize: 15,
    backgroundColor: '#646464',
    color: '#ffffff',
    textAlign: 'center',
  },
  codeline: {
    whiteSpaceTreatment: 'nowrap'
  }
};

export default CodeEditor;