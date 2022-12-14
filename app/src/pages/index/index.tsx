/* eslint-disable jsx-quotes */
import {
  Component,
  PropsWithChildren,
  ReactNode,
  useCallback,
  useEffect,
  useState,
} from "react";
import Taro from "@tarojs/taro";
import { View, Text, ScrollView } from "@tarojs/components";
import { GuideOutlined, SettingOutlined } from "@taroify/icons";

import {
  ActionSheet,
  Avatar,
  Button,
  Dialog,
  Flex,
  Form,
  Loading,
  Textarea,
} from "@taroify/core";
import { IMessage } from "src/types/IMessage";

import "./index.scss";

export interface IMessageProps {
  from: string;
  content: ReactNode | string;
}

function IndexView() {
  const [messages, setMessages] = useState<IMessageProps[]>([]);

  const [loading, setLoading] = useState(false);
  const [actionShow, setActionShow] = useState(false);
  const [question, setQuestion] = useState("");
  const [resetShow, setResetShow] = useState(false);
  const [formBottom, setFormBottom] = useState(30);
  const [firstShow, setFirstShow] = useState(false);

  useEffect(() => {
    Taro.getStorage({
      key: "wx_first_use",
      fail: (err) => {
        setFirstShow(true);
      },
    });
  }, []);

  const socketSubmit = useCallback(
   async (input:string) => {
    setMessages([
      ...messages,
      ...[
        { from: "I", content: input },
        { from: "N", content: <Loading type="spinner" /> },
      ],
    ]);
    Taro.pageScrollTo({ selector: ".scroll-view", offsetTop: 9999 });

    Taro.connectSocket({
      url: `${process.env.WSURL}/ws/v1/wechat/message`,
      success: function () {
        console.log('connect success')
      }
    }).then(task => {
      task.onOpen(function () {
        console.log('onOpen')
        task.send({ data: JSON.stringify({message: input}) })
      })
      task.onMessage(function (resp) {
        const msg = JSON.parse(resp.data)
        console.log('onMessage: ', msg)
        setMessages([
          ...messages,
          ...[
            { from: "I", content: input },
            { from: "N", content: msg.message },
          ],
        ]);
        Taro.pageScrollTo({ selector: ".scroll-view", offsetTop: 9999 });

      })
      task.onError(function () {
        console.log('onError')
      })
      task.onClose(function (e) {
        console.log('onClose: ', e)
      })
    }).finally(() => {
      setLoading(false)
    })
   }, [messages]
  )

  const handleSubmit = useCallback(
    async (input: string) => {
      console.log("b", messages);
      setLoading(true);
      setMessages([
        ...messages,
        ...[
          { from: "I", content: input },
          { from: "N", content: <Loading type="spinner" /> },
        ],
      ]);
      Taro.pageScrollTo({ selector: ".scroll-view", offsetTop: 9999 });
      try {
        let token: IMessage = {} as IMessage;
        try {
          const r = await Taro.getStorage({
            key: "wx_msg_token",
            success: function (cookie) {
              return cookie.data;
            },
            fail: (err) => {
              return "";
            },
          });
          token = r.data;
        } catch (err) {
          console.log("get token", err);
        }
        Taro.pageScrollTo({ selector: ".scroll-view", offsetTop: 9999 });
        console.log("make request");
        const resp = await Taro.request({
          url: `${process.env.URL}/api/v1/wechat/message`,
          method: "POST",
          data: {
            ...token,
            message: input,
          },
          success: (result) => {
            Taro.setStorage({ key: "wx_msg_token", data: result.data });
            return result.data;
          },
        });
        setMessages([
          ...messages,
          { from: "I", content: input },
          { from: "N", content: resp.data.message },
        ]);
        Taro.pageScrollTo({ selector: ".scroll-view", offsetTop: 9999 });
        console.log("a", messages);
      } finally {
        setLoading(false);
      }
    },
    [messages]
  );

  return (
    <View style={{ backgroundColor: "#F7F7F7" }}>
      <ScrollView
        className="scroll-view"
        style={{
          paddingBottom: `${formBottom}px`,
          marginBottom: `${formBottom}px`,
        }}
        enhanced
      >
        {messages !== undefined && messages !== null && messages.length > 0 ? (
          messages.map((m, i) => {
            return (
              <Flex
                gutter={20}
                key={i}
                style={{
                  paddingTop: "10px",
                  paddingBottom: "10px",
                  backgroundColor: m.from === "N" ? "#F7F7F8" : "white",
                }}
              >
                <Flex.Item span={4}>
                  {m.from === "N" ? (
                    <Avatar
                      size="small"
                      shape="square"
                      src="https://i.postimg.cc/NK0dNhSh/apple-touch-icon.png"
                    />
                  ) : (
                    <Avatar
                      style={{ background: "green" }}
                      size="small"
                      shape="square"
                    >
                      {m.from}
                    </Avatar>
                  )}
                </Flex.Item>
                <Flex.Item span={20}>
                  {typeof m.content === "string" ? (
                    <Text onClick={() => Taro.setClipboardData({data: m.content as string})}>{m.content}</Text>
                  ) : (
                    m.content
                  )}
                </Flex.Item>
              </Flex>
            );
          })
        ) : (
          <Flex
            gutter={20}
            style={{
              paddingTop: "10px",
              paddingBottom: "10px",
              backgroundColor: "#F7F7F8",
            }}
          >
            <Flex.Item span={4}>
              <Avatar
                size="small"
                shape="square"
                src="https://i.postimg.cc/NK0dNhSh/apple-touch-icon.png"
              />
            </Flex.Item>
            <Flex.Item span={20}>
              <Text>
                ?????????????????????????????????????????????????????????????????????????????????????????????????????????
                Javascript ????????? HTTP ?????????
              </Text>
            </Flex.Item>
          </Flex>
        )}
        <View style={{ height: "20vh" }}></View>
      </ScrollView>
      <Form
        className="input-tab"
        onSubmit={(e) => {
          setQuestion("");
          socketSubmit(e.detail.value?.message);
        }}
        style={{ paddingBottom: `${formBottom}px` }}
      >
        <Form.Item
          name="message"
          rules={[
            {
              required: true,
              validator: (val) => val !== "",
              message: "????????????",
            },
          ]}
        >
          <Form.Control>
            <Button
              variant="text"
              size="small"
              color="primary"
              icon={<SettingOutlined />}
              onClick={() => setActionShow(true)}
            />
            <Textarea
              onKeyboardHeightChange={(res) => {
                console.log(res);
                setFormBottom(res.detail.height);
              }}
              value={question}
              autoHeight
              placeholder="???????????????"
              onChange={(e) => setQuestion(e.detail.value)}
            />
            <Button
              variant="text"
              size="small"
              color="primary"
              disabled={loading}
              icon={<GuideOutlined />}
              formType="submit"
            />
          </Form.Control>
        </Form.Item>
      </Form>
      <ActionSheet
        open={actionShow}
        onSelect={() => setActionShow(false)}
        onCancel={() => setActionShow(false)}
        onClose={setActionShow}
      >
        <ActionSheet.Action
          value="1"
          name="??????????????????"
          onClick={() => {
            setResetShow(true);
          }}
        />
        <ActionSheet.Action
          value="3"
          name="??????"
          onClick={() => setFirstShow(true)}
        />
        <ActionSheet.Button type="cancel">??????</ActionSheet.Button>
      </ActionSheet>
      <Dialog open={resetShow} onClose={setResetShow}>
        <Dialog.Content>?????????????????????????????????????????????????????????</Dialog.Content>
        <Dialog.Actions>
          <Button
            onClick={() => Taro.redirectTo({ url: "/pages/index/index" })}
          >
            ??????
          </Button>
        </Dialog.Actions>
      </Dialog>
      <Dialog
        open={firstShow}
        onClose={() => {
          setFirstShow(false);
          Taro.setStorage({ key: "wx_first_use", data: "false" });
        }}
      >
        <Dialog.Content>
          <Flex gutter={10} wrap="wrap">
            <Flex.Item span={24}>
              1. ???????????? OpenAI ????????? ChatGPT ??????
            </Flex.Item>
            <Flex.Item span={24}>
              2. ??????????????????????????????????????? proxy ??????
            </Flex.Item>
            <Flex.Item span={24}>
              3. ???????????????????????????????????????????????????????????????????????????
            </Flex.Item>
            <Flex.Item span={24}>
              4. ???????????????????????????????????????????????????
            </Flex.Item>
            <Flex.Item span={24}>
              5. ???????????????????????????????????????????????????????????????
            </Flex.Item>
          </Flex>
        </Dialog.Content>
        <Dialog.Actions>
          <Button
            onClick={() => {
              setFirstShow(false);
              Taro.setStorage({ key: "wx_first_use", data: "false" });
            }}
          >
            ??????
          </Button>
        </Dialog.Actions>
      </Dialog>
    </View>
  );
}

export default class Index extends Component<PropsWithChildren> {
  componentWillMount() {}

  componentDidMount() {}

  componentWillUnmount() {}

  componentDidShow() {}

  componentDidHide() {}

  render() {
    Taro.pageScrollTo({ selector: ".scroll-view", scrollTop: 9999 });
    return (
      <View className="index">
        <IndexView />
      </View>
    );
  }
}
