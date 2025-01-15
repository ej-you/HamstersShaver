import json
import os
from asyncio import Lock
from typing import Optional, Dict

import asyncio
import aiofiles

from tonutils.tonconnect import TonConnect, IStorage
from tonutils.tonconnect.models import Event, EventError, SendTransactionResponse
from tonutils.tonconnect.utils.exceptions import TonConnectError


class FileStorage(IStorage):

    def __init__(self, file_path: str):
        self.file_path = file_path
        self.lock = Lock()

        if not os.path.exists(self.file_path):
            with open(self.file_path, 'w') as f:
                json.dump({}, f)  # type: ignore

    async def _read_data(self) -> Dict[str, str]:
        async with self.lock:
            async with aiofiles.open(self.file_path, 'r') as f:
                content = await f.read()
                if content:
                    return json.loads(content)
                return {}

    async def _write_data(self, data: Dict[str, str]) -> None:
        async with self.lock:
            async with aiofiles.open(self.file_path, 'w') as f:
                await f.write(json.dumps(data, indent=4))

    async def set_item(self, key: str, value: str) -> None:
        data = await self._read_data()
        data[key] = value
        await self._write_data(data)

    async def get_item(self, key: str, default_value: Optional[str] = None) -> Optional[str]:
        data = await self._read_data()
        return data.get(key, default_value)

    async def remove_item(self, key: str) -> None:
        data = await self._read_data()
        if key in data:
            del data[key]
            await self._write_data(data)


TC_MANIFEST_URL = "https://raw.githubusercontent.com/nessshon/tonutils/main/examples/tonconnect/tonconnect-manifest.json"  # noqa
# In this example, FileStorage from storage.py is used
TC_STORAGE = FileStorage("connection.json")

USER_ID = 123
RPC_REQUEST_ID = 123

WALLET_ADDRESS = "UQDLjMTSj58tHMc9oIeo00dbV6smPp1plQw5vUSZADeUInM9"


# Create an instance of TonConnect with the specified storage and manifest
tc = TonConnect(storage=TC_STORAGE, manifest_url=TC_MANIFEST_URL, include_wallets=[WALLET_ADDRESS])


@tc.on_event(Event.TRANSACTION)
async def on_transaction(transaction: SendTransactionResponse) -> None:
    print(f"[Transaction SENT] Transaction successfully sent. Message hash: {transaction.hash}")


@tc.on_event(EventError.TRANSACTION)
async def on_transaction_error(error: TonConnectError) -> None:
    print(f"[Transaction ERROR] Transaction failed. Error: {error}")


async def main() -> None:
    # Initialize the connector for the user
    connector = await tc.init_connector(USER_ID)
    print("connector:", connector)

    print(await tc.get_wallets())

    # async with connector.pending_transaction_context(RPC_REQUEST_ID) as response:
    #     if isinstance(response, TonConnectError):
    #         print(f"Error sending transaction: {response.message}")
    #     else:
    #         print(f"Transaction successful! Hash: {response.hash}")

    # Close all TonConnect connections
    await tc.close_all()


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except (KeyboardInterrupt, SystemExit):
        # Ensure all connections are closed in case of interruption
        asyncio.run(tc.close_all())
