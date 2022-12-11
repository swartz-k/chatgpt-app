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
                你可以问我：用几句话介绍一下量子计算？有没有关于生日的创意？或者如何用
                Javascript 写一个 HTTP 请求？
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
              message: "不能为空",
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
              placeholder="说点什么吧"
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
          name="重新开始会话"
          onClick={() => {
            setResetShow(true);
          }}
        />
        <ActionSheet.Action
          value="3"
          name="关于"
          onClick={() => setFirstShow(true)}
        />
        <ActionSheet.Button type="cancel">取消</ActionSheet.Button>
      </ActionSheet>
      <Dialog open={resetShow} onClose={setResetShow}>
        <Dialog.Content>服务不会保存对话记录，重启后会消失哦～</Dialog.Content>
        <Dialog.Actions>
          <Button
            onClick={() => Taro.redirectTo({ url: "/pages/index/index" })}
          >
            确认
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
              1. 欢迎体验 OpenAI 开放的 ChatGPT 模型
            </Flex.Item>
            <Flex.Item span={24}>
              2. 本服务不保存会话记录，只做 proxy 转发
            </Flex.Item>
            <Flex.Item span={24}>
              3. 有趣的记录，可以截图保存，重新进入后对话记录将消失
            </Flex.Item>
            <Flex.Item span={24}>
              4. 会话过程稍慢，请耐心，程序猿优化中
            </Flex.Item>
            <Flex.Item span={24}>
              5. 任何疑问，欢迎打开小程序资料，联系客服反馈
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
            确认
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
