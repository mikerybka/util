import React, { useEffect, useState } from "react";
import sendLoginCode from "../functions/sendLoginCode";
import login from "../functions/login";
import Form from "./Form";
import PhoneField from "./PhoneField";
import CodeField from "./CodeField";
import SuccessPrompt from "./SuccessPrompt";

export default function Login() {
    const [stage, setStage] = useState<"phone" | "code" | "success">("phone");
    const [phone, setPhone] = useState("");
    const [code, setCode] = useState("");

    const submitPhone = async () => {
        await sendLoginCode(phone);
    }

    const submitCode = async () => {
        return await login(phone, code)
    }

    useEffect(() => {
        if (stage !== "success") {
            const session = localStorage.getItem('session');
            if (session) {
                setStage("success")
            }
        }
    }, [])

    if (stage === "phone") {
        return <Form onSubmit={submitPhone} onSuccess={() => {
            setStage("code");
        }}>
            <PhoneField id="phone" name="Phone" value={phone} onChange={setPhone} />
        </Form>
    }

    if (stage === "code") {
        return <Form onSubmit={submitCode} onSuccess={(res) => {
            localStorage.setItem('session', JSON.stringify(res))
            setStage("success");
        }}>
            <CodeField digits={6} id="code" name="Code" value={code} onChange={setCode} />
        </Form>
    }

    if (stage === "success") {
        return <SuccessPrompt message="You have logged in successfully." actions={[
            {
                primary: true,
                text: "Schemas",
                href: "/" + phone,
            },
            {
                text: "Logout",
                onClick: async () => {
                    localStorage.clear();
                    window.location.href = "/";
                },
            }
        ]} />
    }
}
