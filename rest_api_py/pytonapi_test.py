import os

from pytonapi import AsyncTonapi
from pytonapi.schema.events import TransactionEventData

# Enter your API key
API_KEY = os.getenv("TON_API_KEY")

SUCCESS_TRANS_HASH = "99821a8101a7d25e76811e01e97d606d5caa5f62419867245b8fd1f5f362590b"
FAILED_TRANS_HASH = "9225f9db67e9857e177c1e3570de26c157e94feed7d989f63244972879833a0a"

# List of TON blockchain accounts to monitor
ACCOUNTS = [
    "UQDLjMTSj58tHMc9oIeo00dbV6smPp1plQw5vUSZADeUInM9",

    "UQCJYWj8oz1sceKGTqLELUhJ9jAJynXRiRW2AqDpTb-8utwN",
    "UQDJmlxQ6Bd7N_tmHuf2p-imf8H5CwIunUSJI5C15H5gUqTY",
    "UQADRQ7WjDlbuNLw7jWKSWE7Si_oA1t_JWDyXtotfCyaP0qY",
    "UQCYbWPeHQSw4lwUBIwcnxR6avVwIpax4f-xR-ba2N6A0CEF",
    "EQCuqtljI1F4CILgj9o-th3vqB9eNF_mt3jtyDIEWz--pojh",
    "UQDEd5vftG87eBetMbGrb2BQdKgj0bgkJqMm6IymmVeIlTtM",
    "UQC9yHiF6CtrtwLpQZfBzq8YiLUrczBPq4IQt5H8tili-gbD",
    "UQDDFR4bnkDaKoFk26gaOkxfnsFlhtt4XidUm81xfieUHIpj",
    "UQAtLX1jSk459g5y7pPqw4SkEYMnDZgxd7nG-iMBUVqwaN--",
    "UQAZgA1LHMrIiYiGmaOGSujQTasommm7LFRSrThtElr9IZuZ",
    "UQDPy5vjHGOLvEBlHpiFZgbU7jHpNoLU-FNW-ZkUcLbEb0FF",
    "UQA4mfrV45OEIuTyJKDQe41FX1X0XD8IPJ9UYb7Tpu3gK6kO",  # какой-то бот, поэтому очень хорошо его отслеживать
]


async def handler(event: TransactionEventData, tonapi: AsyncTonapi) -> None:
    """
    Handle SSEvent for transactions.

    :param event: The SSEvent object containing transaction details.
    :param tonapi: An instance of AsyncTonapi for interacting with TON API.
    """
    print("event:", event)

    trace = await tonapi.traces.get_trace(event.tx_hash)
    print("trace:", trace)


async def main() -> None:
    tonapi = AsyncTonapi(api_key=API_KEY)
    print("tonapi:", tonapi)

    # Subscribe to transaction events for the specified accounts
    # await tonapi.sse.subscribe_to_transactions(
    #     handler, accounts=ACCOUNTS, args=(tonapi,)
    # )

    # trans_hash = "837edbb2f337f27f1e25c7d87299d468303656e716f051e8b4ad4301d5e0c8b3"
    trans_trace = await tonapi.traces.get_trace(FAILED_TRANS_HASH)
    # print("one trace:", trace)

    def trace_unwrap(trace):
        print("trace:", trace.transaction)
        print()
        if trace.children:
            trace_unwrap(trace.children[0])

    trace_unwrap(trans_trace)


async def main2() -> None:
    tonapi = AsyncTonapi(api_key=API_KEY)
    print("tonapi:", tonapi)

    trans_data = await tonapi.blockchain.get_transaction_data(transaction_id=FAILED_TRANS_HASH)
    print("trans_data.success:", trans_data.success)
    print("trans_data:", trans_data)


if __name__ == "__main__":
    import asyncio

    asyncio.run(main())
